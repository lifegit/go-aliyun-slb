package app

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func JsonError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": msg})
}
func JsonErrorData(c *gin.Context, data interface{}, msg string) {
	c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": msg, "data": data})
}
func JsonData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func JsonSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func JsonErrorMapData(c *gin.Context, data string, err error) bool {
	if err != nil {
		m := map[string]interface{}{}
		json.Unmarshal([]byte(data), &m)
		JsonErrorData(c, m, "失败")
		return true
	}
	return false
}

func HandleError(c *gin.Context, err error) bool {
	if err != nil {
		JsonError(c, err.Error())
		return true
	}
	return false
}

func HandleErrorData(c *gin.Context, data interface{}, err error) bool {
	if err != nil {
		JsonErrorData(c, data, err.Error())
		return true
	}
	return false
}

func ParseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return uint(parseId), nil
}
