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
 2.

 ```