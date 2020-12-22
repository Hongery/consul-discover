### 微服务

Consul服务、Go-kit微服务工具

下载 consul 工具https://www.consul.io/downloads.html

```bash
//linux环境
hongery@LAPTOP-52395K7R:/mnt/c/Users/ZYKJ/Downloads/consul_linux$ ./consul  version
Consul v1.9.0
Revision a417fe510
Protocol 2 spoken by default, understands 2 to 3 (agent will automatically use protocol >2 when speaking to compatible agents)

//任意路径都能使用Consul命令
hongery@LAPTOP-52395K7R:/mnt/c/Users/ZYKJ/Downloads/consul_linux$ sudo mv consul /usr/local/bin/
[sudo] password for hongery: 

//启动Consul
hongery@LAPTOP-52395K7R:/usr/local/bin$ consul agent -dev

//浏览器查看
http://localhost:8500/
```

Go-kit工具

```
go get github.com/go-kit/kit
```

#### 1. 项目目录结构

```bash
config      //config配置   
discover    //服务发现客户端
endpoint    //用于接受请求并返回响应,将请求转化为sercice接口可以处理的参数，并返回给transport层
service  	//业务代码实现，服务注册和发现相关实现
transport   //项目提供的服务方式   HTTP服务
main.go 
README.md  
```
#### 2.main.go

```
1.声明并初始化DiscoveryClient，调用Register方法完成服务注册
2.生命并初始化接口service，并基于service构建Endpoint，接着构建对应的http.Handle，对外暴露http接口，启动http服务
3.注册关闭事件，检测服务关闭事件
```

#### http 方式交互（函数执行）
```
 1.先启动consul
 2.go run main.go （服务注册）
    hongery@LAPTOP-52395K7R:/mnt/d/gocode/src/github.com/consul-discover$ go run main.go 
    2020/12/22 17:09:11 Http Server start at port:10086
    2020/12/22 17:09:11 Register Service Success!
 3.登入http://localhost:8500/ui/dc1/services/查看
 4.服务注销（ctrl+c）
    ^C 2020/12/22 17:13:28 Deregister Service Success!
    2020/12/22 17:13:28 interrupt
 5.服务发现   
    go run main.go (服务得运行)
    浏览器运行http://127.0.0.1:10086/discovery?serviceName=SayHello
    结果：
        {"instances":[{"ID":"SayHello-af4dd723-a44a-4f82-aac2-5b7f58bc0f4f","Service":"SayHello","Tags":[],"Meta":null,"Port":10086,"Address":"127.0.0.1","TaggedAddresses":{"lan_ipv4":{"Address":"127.0.0.1","Port":10086},"wan_ipv4":{"Address":"127.0.0.1","Port":10086}},"Weights":{"Passing":1,"Warning":1},"EnableTagOverride":false,"CreateIndex":408,"ModifyIndex":408,"Proxy":{"MeshGateway":{},"Expose":{}},"Connect":{}}],"error":""}

 ```

 #### Go-kit 服务注册与发现包与Consul交互
 在目录discover/kit_discover_client.go

 ``go get github.com/hashicorp/consul``
 
 ``go get github.com/go-kit/kit ``

 ```
 1.go mod tidy 引入外部依赖

 2.需要替换（main.go）
   //discoveryClient,err := discover.NewMyDiscoverClient(*consulHost,*consulPort)
	discoveryClient, err := discover.NewKitDiscoverClient(*consulHost, *consulPort)

 ```


 #### 提交代码
 ```bash
PS D:\gocode\src\github.com\consul-discover> git init
Initialized empty Git repository in D:/gocode/src/github.com/consul-discover/.git/
PS D:\gocode\src\github.com\consul-discover> git add .
PS D:\gocode\src\github.com\consul-discover> git commit -m "first commit "
[master (root-commit) e8edb51] first commit
 11 files changed, 1623 insertions(+)
 create mode 100644 README.md
 create mode 100644 config/config.go
 create mode 100644 discover/discover_client.go
 create mode 100644 discover/kit_discover_client.go
 create mode 100644 endpoint/endpoints.go
 create mode 100644 go.mod
 create mode 100644 go.sum
 create mode 100644 main.go
 create mode 100644 service/service.go
 create mode 100644 transport/http.go
PS D:\gocode\src\github.com\consul-discover> git branch -M main (main分支)
PS D:\gocode\src\github.com\consul-discover> git remote add origin  https://github.com/Hongery/consul-discover.git
PS D:\gocode\src\github.com\consul-discover> git push origin main
Enumerating objects: 18, done.
Counting objects: 100% (18/18), done.
Delta compression using up to 8 threads
Compressing objects: 100% (14/14), done.
Writing objects: 100% (18/18), 42.81 KiB | 7.13 MiB/s, done.
Total 18 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/Hongery/consul-discover.git
 * [new branch]      main -> main

 ```
 
 
 #### my_discover_client 和 kit_discover_client 的区别
 ```
 my_discover_client 配置信息需要自己手动配置，需要的信息自己添加
 kit_discover_client 配置信息已经完全集成
 ```
 MyDiscoverClient(*consulHost,*consulPort)
	discoveryClient, err := discover.NewKitDiscoverClient(*consulHost, *consulPort)
