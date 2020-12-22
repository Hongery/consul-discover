package endpoint

import (
	"context"

	"github.com/consul-discover/service"
	"github.com/go-kit/kit/endpoint"
)

type DiscoveryEndpoints struct {
	SayHelloEndpoint    endpoint.Endpoint //go-kit/kit/endpoint
	DiscoveryEndpoint   endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

// 打招呼请求结构体
type SayHelloRequest struct {
}

// 打招呼响应结构体
type SayHelloResponse struct {
	Message string `json:"message"`
}

// 创建打招呼 Endpoint  请求
// Endpoint类型是一个函数 //  type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
func MakeSayHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		message := svc.SayHello()
		return SayHelloResponse{
			Message: message,
		}, nil
	}
}

// 服务发现请求结构体
type DiscoveryRequest struct {
	ServiceName string
}

// 服务发现响应结构体
type DiscoveryResponse struct {
	Instances []interface{} `json:"instances"`
	Error     string        `json:"error"`
}

// 创建服务发现的 Endpoint
func MakeDiscoveryEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DiscoveryRequest)
		instances, err := svc.DiscoveryService(ctx, req.ServiceName) //调用服务发现service
		var errString = ""
		if err != nil {
			errString = err.Error()
		}
		return &DiscoveryResponse{
			Instances: instances,
			Error:     errString,
		}, nil
	}
}

// HealthRequest 健康检查请求结构
type HealthRequest struct{}

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status bool `json:"status"`
}

// MakeHealthCheckEndpoint 创建健康检查Endpoint
func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()  //调用service 健康检查
		return HealthResponse{
			Status: status,
		}, nil
	}
}
