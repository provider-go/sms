package district

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/types"
	"github.com/provider-go/sms/global"
	"github.com/provider-go/sms/router"
)

type Plugin struct{}

func CreatePlugin() *Plugin {
	return &Plugin{}
}

func CreatePluginAndDB(instance types.PluginNeedInstance) *Plugin {
	global.DB = instance.Mysql
	global.Cache = instance.Cache
	global.SMCC = instance.SMCC
	return &Plugin{}
}

func (*Plugin) Register(group *gin.RouterGroup) {
	router.GroupApp.InitRouter(group)
}

func (*Plugin) RouterPath() string {
	return "sms"
}
