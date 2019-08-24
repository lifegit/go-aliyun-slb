# go-aliyun-slb

#### 介绍
使用阿里云SDK对SLB(负载均衡)进行创建，添加监听端口，释放

#### 前提说明

conf/base.toml 配置好相关信息

#### 使用说明

流程:

    1.创建slb
    2.添加后端服务器
    3.创建监听
    4.开始监听
    5.删除slb

#### API
##### SLB:

    创建slb:
    http://localhost:5555/api/slb/CreateLoadBalancer
    参数
        internetChargeType   计费方式 eg: 按带宽计费:paybybandwidth  按流量计费:paybytraffic
        bandwidth            带宽峰值,单位:M

    添加后端服务器
    http://localhost:5555/api/slb/AddBackendServers?slbId=?&ecsId=?
    参数
        slbId   刚创建的slb的id
        ecsId   要添加的ecs服务器的id(阿里云规定slb必须和ecs在同一个地域)


    创建监听
    http://localhost:5555/api/slb/CreateLoadBalancerTCPListener?slbId=?&listenerPort=?&backendServerPort=?
    参数
        slbId               刚创建的slb的id
        listenerPort        slb开放端口 
        backendServerPort   ecs开放端口


    开始监听
    http://localhost:5555/api/slb/StartLoadBalancerListener?slbId=?&listenerPort=?
    参数
        slbId               刚创建的slb的id
        listenerPort        slb开放端口 


    删除slb
    http://localhost:5555/api/slb/DeleteLoadBalancer?slbId=?
    参数
        slbId               要删除的slb的id
     
##### DNS   
    修改解析记录
    http://localhost:5555/api/dns/UpdateDomainRecord?recordId=?&rr=?&type=?&value=?
    参数
        recordId 解析记录的ID
        rr      主机记录,如果要解析@.exmaple.com，主机记录要填写”@”。
        type    解析记录类型，例:A
        value   记录值
        ttl     解析时间