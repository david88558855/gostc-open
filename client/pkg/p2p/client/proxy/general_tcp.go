// Copyright 2023 The frp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	v1 "gostc-sub/pkg/p2p/pkg/config/v1"
	"reflect"
)

func init() {
	pxyConfs := []v1.ProxyConfigurer{
		&v1.STCPProxyConfig{},
	}
	for _, cfg := range pxyConfs {
		RegisterProxyFactory(reflect.TypeOf(cfg), NewGeneralTCPProxy)
	}
}

// GeneralTCPProxy is a general implementation of Proxy interface for TCP protocol.
// If the default GeneralTCPProxy cannot meet the requirements, you can customize
// the implementation of the Proxy interface.
type GeneralTCPProxy struct {
	*BaseProxy
}

func NewGeneralTCPProxy(baseProxy *BaseProxy, _ v1.ProxyConfigurer) Proxy {
	return &GeneralTCPProxy{
		BaseProxy: baseProxy,
	}
}
