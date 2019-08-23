package handlers

import (
	"github.com/gin-gonic/gin"
	"go-aliyun-slb/pkg/app"
	"go-aliyun-slb/pkg/slb"
)

func init() {
	group := groupApi.Group("slb")
	group.GET("CreateLoadBalancer", CreateLoadBalancer)
	group.GET("AddBackendServers", AddBackendServers)
	group.GET("CreateLoadBalancerTCPListener", CreateLoadBalancerTCPListener)
	group.GET("StartLoadBalancerListener", StartLoadBalancerListener)
	group.GET("DeleteLoadBalancer", DeleteLoadBalancer)
}

func CreateLoadBalancer(c *gin.Context) {
	res, err := slb.CreateLoadBalancer()
	if app.JsonErrorMapData(c, res.GetHttpContentString(), err) {
		return
	}

	app.JsonData(c, res)
}

func AddBackendServers(c *gin.Context) {
	// 接收参数
	var param struct {
		SlbId string `binding:"slbId"`
		EcsId string `binding:"ecsId"`
	}
	err := c.ShouldBind(&param)
	if app.HandleError(c, err) {
		return
	}

	res, err := slb.AddBackendServers(param.SlbId, param.EcsId)
	if app.JsonErrorMapData(c, res.GetHttpContentString(), err) {
		return
	}

	app.JsonData(c, res)
}

func CreateLoadBalancerTCPListener(c *gin.Context) {
	// 接收参数
	var param struct {
		SlbId             string `binding:"slbId"`
		ListenerPort      int    `binding:"listenerPort"`
		BackendServerPort int    `binding:"backendServerPort"`
	}
	err := c.ShouldBind(&param)
	if app.HandleError(c, err) {
		return
	}

	res, err := slb.CreateLoadBalancerTCPListener(param.SlbId, param.ListenerPort, param.BackendServerPort)
	if app.JsonErrorMapData(c, res.GetHttpContentString(), err) {
		return
	}

	app.JsonData(c, res)
}

func StartLoadBalancerListener(c *gin.Context) {
	// 接收参数
	var param struct {
		SlbId        string `binding:"slbId"`
		ListenerPort int    `binding:"listenerPort"`
	}
	err := c.ShouldBind(&param)
	if app.HandleError(c, err) {
		return
	}

	res, err := slb.StartLoadBalancerListener(param.SlbId, param.ListenerPort)
	if app.JsonErrorMapData(c, res.GetHttpContentString(), err) {
		return
	}

	app.JsonData(c, res)
}

func DeleteLoadBalancer(c *gin.Context) {
	// 接收参数
	var param struct {
		SlbId string `binding:"slbId"`
	}
	err := c.ShouldBind(&param)
	if app.HandleError(c, err) {
		return
	}

	res, err := slb.DeleteLoadBalancer(param.SlbId)
	if app.JsonErrorMapData(c, res.GetHttpContentString(), err) {
		return
	}

	app.JsonData(c, res)
}
