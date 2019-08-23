package slb

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"go-aliyun-slb/pkg/conf"
	"go-aliyun-slb/pkg/logging/normalLogging"
)

/**
流程：
1.创建slb
2.添加后端服务器
3.创建监听
4.开始监听
*/
var client *slb.Client

func init() {
	c, err := slb.NewClientWithAccessKey(conf.GetString("aliyun.slb.regin_id"), conf.GetString("aliyun.access_key_id"), conf.GetString("aliyun.access_key_secret"))
	if err != nil {
		normalLogging.Logger.WithError(err).Fatal(err)
		return
	}
	client = c
}

/**
删除slb
*/
func DeleteLoadBalancer(balancerId string) (response *slb.DeleteLoadBalancerResponse, err error) {
	request := slb.CreateDeleteLoadBalancerRequest()
	request.Scheme = "https"
	request.LoadBalancerId = balancerId

	response, err = client.DeleteLoadBalancer(request)

	return
}

/**
创建监听
*/
func CreateLoadBalancerTCPListener(balancerId string, listenerPort int, backendServerPort int) (response *slb.CreateLoadBalancerTCPListenerResponse, err error) {
	request := slb.CreateCreateLoadBalancerTCPListenerRequest()
	request.Scheme = "https"
	request.LoadBalancerId = balancerId
	request.Bandwidth = requests.NewInteger(-1)
	request.ListenerPort = requests.NewInteger(listenerPort)
	request.BackendServerPort = requests.NewInteger(backendServerPort)

	response, err = client.CreateLoadBalancerTCPListener(request)

	return
}

/**
开始监听
*/
func StartLoadBalancerListener(balancerId string, listenerPort int) (response *slb.StartLoadBalancerListenerResponse, err error) {
	request := slb.CreateStartLoadBalancerListenerRequest()
	request.Scheme = "https"
	request.LoadBalancerId = balancerId
	request.ListenerPort = requests.NewInteger(listenerPort)

	response, err = client.StartLoadBalancerListener(request)

	return
}

/**
创建slb
*/
func CreateLoadBalancer() (response *slb.CreateLoadBalancerResponse, err error) {
	request := slb.CreateCreateLoadBalancerRequest()
	request.Scheme = "https"

	response, err = client.CreateLoadBalancer(request)

	return
}

/**
添加后端服务器
*/
func AddBackendServers(balancerId string, serverId string) (response *slb.AddBackendServersResponse, err error) {
	request := slb.CreateAddBackendServersRequest()
	request.Scheme = "https"

	request.LoadBalancerId = balancerId
	request.BackendServers = `[{ "ServerId": "` + serverId + `", "Weight": "100", "Type": "ecs"}]`
	fmt.Println(request.BackendServers)
	response, err = client.AddBackendServers(request)

	return
}
