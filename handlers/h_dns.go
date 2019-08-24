package handlers

import (
	"github.com/gin-gonic/gin"
	"go-aliyun-slb/pkg/app"
	"go-aliyun-slb/pkg/dns"
)

func init() {
	group := groupApi.Group("dns")
	group.GET("UpdateDomainRecord", UpdateDomainRecord)
}

func UpdateDomainRecord(c *gin.Context) {
	// 接收参数
	var param struct {
		RecordId string `form:"recordId" binding:"required"`
		RR       string `form:"rr" binding:"required"`
		Type     string `form:"type" binding:"required"`
		Value    string `form:"value" binding:"required"`
		Ttl      int    `form:"ttl" binding:"required"`
	}
	err := c.ShouldBind(&param)
	if app.HandleError(c, err) {
		return
	}

	res, err := dns.UpdateDomainRecord(param.RecordId, param.RR, param.Type, param.Value, param.Ttl)
	if app.JsonErrorMapData(c, res.GetHttpContentString(), err) {
		return
	}

	app.JsonData(c, res)
}
