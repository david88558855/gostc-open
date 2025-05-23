package gost_engine

import (
	"encoding/json"
	"errors"
	"github.com/lxzan/gws"
	"go.uber.org/zap"
	"server/global"
	"server/model"
	"server/repository"
	"server/service/common/cache"
	"sync/atomic"
	"time"
)

type NodeEvent struct {
	code            string
	log             *zap.Logger
	isRunning       bool
	sendAtomic      atomic.Int32
	sendOperationId string
	isRegister      bool

	maxRetries    int // 最大重试次数
	retryInterval time.Duration
	checkInterval time.Duration // 循环检查消息等待间隔
}

func NewNodeEvent(code string, log *zap.Logger) *NodeEvent {
	return &NodeEvent{
		code:          code,
		log:           log,
		maxRetries:    10,
		retryInterval: time.Second,
		checkInterval: time.Second,
	}
}

func (event *NodeEvent) OnOpen(socket *gws.Conn) {
	event.isRunning = true
	if event.code == "" {
		event.isRunning = false
		return
	}
	cache.SetNodeOnline(event.code, true, time.Second*30)
	go func() {
		for {
			time.Sleep(time.Second * 10)
			if err := socket.WritePing(nil); err != nil {
				return
			}
		}
	}()
	go event.sendLoop(socket)
	go event.checkRegister()
}

func (event *NodeEvent) sendLoop(socket *gws.Conn) {
	ticker := time.NewTicker(event.checkInterval)
	defer func() {
		event.log.Info("stop sendLoop",
			zap.String("type", "client"),
			zap.String("code", event.code))
		ticker.Stop()
	}()

	for event.isRunning {
		// Skip if sending is blocked
		if event.sendAtomic.Load() == 1 {
			time.Sleep(event.retryInterval)
			continue
		}

		msgChan, err := msgRegistry.PullMessage(event.code)
		if err != nil {
			time.Sleep(event.retryInterval)
			continue
		}

		select {
		case req := <-msgChan:
			event.handleMessage(socket, req)
		case <-ticker.C:
			// Continue loop
		}
	}
}

func (event *NodeEvent) handleMessage(socket *gws.Conn, req MessageData) {
	event.sendAtomic.Store(1)
	event.sendOperationId = req.OperationId
	event.WriteAny(socket, req)

	for retry := 0; retry < event.maxRetries && event.isRunning; retry++ {
		if event.sendAtomic.Load() == 0 {
			return
		}

		time.Sleep(event.retryInterval)
	}

	// Final attempt if still needed
	if event.isRunning && event.sendAtomic.Load() == 1 {
		event.WriteAny(socket, req)
	}
}

func (event *NodeEvent) register() {
	event.isRegister = true
	db, _, _ := repository.Get("")
	NodeConfig(db, event.code)
}

func (event *NodeEvent) checkRegister() {
	time.Sleep(time.Second * 10)
	if event.isRegister {
		return
	}
	if !event.isRunning {
		return
	}
	event.log.Warn("节点连接注册超时，关闭连接", zap.String("type", "node"), zap.String("code", event.code))
	NodeStop(event.code, "节点连接注册超时")
}

func (event *NodeEvent) OnClose(socket *gws.Conn, err error) {
	if !event.isRunning {
		return
	}
	event.isRunning = false
	event.log.Info("connect close")
	if err != nil && !errors.As(err, &gws.ErrConnClosed) {
		event.log.Info("connect close fail", zap.Error(err))
	}
	cache.SetNodeOnline(event.code, false, time.Second*30)
	msgRegistry.CleanMessage(event.code)
	event.sendAtomic.Store(0)
}

func (event *NodeEvent) OnPing(socket *gws.Conn, payload []byte) {
	cache.SetNodeOnline(event.code, true, time.Second*30)
	_ = socket.SetDeadline(time.Now().Add(time.Second * 30))
}

func (event *NodeEvent) OnPong(socket *gws.Conn, payload []byte) {
}

func (event *NodeEvent) OnMessage(socket *gws.Conn, message *gws.Message) {
	var msg MessageData
	_ = json.Unmarshal(message.Bytes(), &msg)
	if msg.OperationId == event.sendOperationId {
		event.sendAtomic.Store(0)
	}
	switch msg.OperationType {
	case "register":
		var data = make(map[string]any)
		_ = msg.GetContent(&data)
		cache.SetNodeVersion(event.code, data["version"].(string))

		_, customDomain := data["custom_domain"]
		cache.SetNodeCustomDomain(event.code, customDomain)

		event.register()
	case "logger":
		var cfg model.SystemConfigGost
		cache.GetSystemConfigGost(&cfg)
		if cfg.Logger != "1" {
			break
		}
		var logMsg = ""
		_ = msg.GetContent(&logMsg)
		var log = struct {
			Level string `json:"level"`
		}{}
		_ = json.Unmarshal([]byte(logMsg), &log)
		global.DB.GetDB().Create(&model.GostNodeLogger{
			Level:     log.Level,
			NodeCode:  event.code,
			Content:   logMsg,
			CreatedAt: time.Now().Unix(),
		})
	case "port_check":
		var data NodePortCheckResult
		_ = msg.GetContent(&data)
		if data.Code == "" {
			return
		}
		cache.SetNodePortUse(data.Code, data.Port, data.Use, time.Minute)
	}
}

type NodePortCheckResult struct {
	Code string `json:"code"` // 节点编号
	Use  bool   `json:"use"`  // 1=被占用
	Port string `json:"port"` // 端口
}

func (event *NodeEvent) WriteAny(socket *gws.Conn, data any) {
	marshal, _ := json.Marshal(data)
	_ = socket.WriteString(string(marshal))
}
