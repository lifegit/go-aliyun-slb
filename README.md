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
    创建slb:
    http://localhost:5555/api/slb/CreateLoadBalancer


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