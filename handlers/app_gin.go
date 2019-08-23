package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-aliyun-slb/pkg/conf"
	"go-aliyun-slb/pkg/logging/normalLogging"
	"go-aliyun-slb/pkg/paramValidator"
)

var r *gin.Engine
var groupApi *gin.RouterGroup

//in the same package init executes in file'name alphabet order
func init() {
	// logging
	normalLogging.Setup(conf.GetString("log.baseDir"), nil)

	// paramValidator
	paramValidator.Setup()

	r = gin.Default()
	api := "api"
	groupApi = r.Group(api)
}

//ServerRun start the server server
func ServerRun() {
	addr := fmt.Sprintf(":%d", conf.GetInt("app.port"))
	r.Run(addr)
}
