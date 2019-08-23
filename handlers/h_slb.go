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
	// 接收参数
	var param struct {
		InternetChargeType string `form:"internetChargeType" binding:"required,eq=paybybandwidth|eq=paybytraffic`
		Bandwidth          int    `form:"bandwidth" binding:"required,numeric`
	}
	res, err := slb.CreateLoadBalancer(param.InternetChargeType, param.Bandwidth)
	if app.JsonErrorMapData(c, res.GetHttpContentString(), err) {
		return
	}

	app.JsonData(c, res)
}

func AddBackendServers(c *gin.Context) {
	// 接收参数
	var param struct {
		SlbId string `form:"slbId" binding:"required"`
		EcsId string `form:"ecsId" binding:"required"`
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
		SlbId             string `form:"slbId" binding:"required"`
		ListenerPort      int    `form:"listenerPort" binding:"required"`
		BackendServerPort int    `form:"backendServerPort" binding:"required"`
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
		SlbId        string `form:"slbId" binding:"required"`
		ListenerPort int    `form:"listenerPort" binding:"required"`
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
		SlbId string `form:"slbId" binding:"required"`
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
