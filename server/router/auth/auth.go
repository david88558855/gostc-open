package auth

import (
	"github.com/gin-gonic/gin"
	"server/controller/auth"
	"server/global"
	"server/router/middleware"
)

func InitAuth(group *gin.RouterGroup) {
	group.POST("login", auth.Login)
	group.POST("register", auth.Register)
	group.POST("loginOtp", auth.LoginOtp)
	group.POST("captcha", middleware.IpLimiter, auth.Captcha)
	group.POST("renew", middleware.Auth(global.Jwt), auth.Renew)
	group.POST("password", middleware.Auth(global.Jwt), auth.Password)
	group.POST("userInfo", middleware.Auth(global.Jwt), auth.UserInfo)
	group.POST("checkin", middleware.Auth(global.Jwt), auth.Checkin)
	group.POST("genOtp", middleware.Auth(global.Jwt), auth.GenOtp)
	group.POST("openOtp", middleware.Auth(global.Jwt), auth.OpenOtp)
	group.POST("closeOtp", middleware.Auth(global.Jwt), auth.CloseOtp)
	group.POST("bindEmail", middleware.Auth(global.Jwt), auth.BindEmail)
	group.POST("genBindEmailCode", middleware.Auth(global.Jwt), auth.GenBindEmailCode)
	group.POST("unBindEmail", middleware.Auth(global.Jwt), auth.UnBindEmail)
	group.POST("genResetPwdEmailCode", auth.GenResetPwdEmailCode)
	group.POST("resetPwd", auth.ResetPwd)
}
