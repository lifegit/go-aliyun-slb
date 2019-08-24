package dns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"go-aliyun-slb/pkg/conf"
	"go-aliyun-slb/pkg/logging/normalLogging"
)

var client *alidns.Client

func init() {
	c, err := alidns.NewClientWithAccessKey(conf.GetString("aliyun.slb.regin_id"), conf.GetString("aliyun.access_key_id"), conf.GetString("aliyun.access_key_secret"))
	if err != nil {
		normalLogging.Logger.WithError(err).Fatal(err)
		return
	}
	client = c
}

/**
更新DNS解析
*/
func UpdateDomainRecord(recordId, rr, typee, value string, ttl int) (response *alidns.UpdateDomainRecordResponse, err error) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = recordId
	request.RR = rr
	request.Type = typee
	request.Value = value
	request.TTL = requests.NewInteger(ttl)

	response, err = client.UpdateDomainRecord(request)

	return
}
