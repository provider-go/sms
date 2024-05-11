package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/sms/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		// sms_logs 表操作
		Router.POST("sendByAli", api.SendCodeByAli)
		Router.POST("sendBySandbox", api.SendCodeBySandbox)
		Router.GET("logList", api.LogList)
	}
}
