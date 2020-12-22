package service

import (
	"context"
	"errors"

	"github.com/consul-discover/discover"
	"github.com/consul-discover/config"
)

type Service interface {

	// HealthCheck check service health status
	HealthCheck() bool

	// sayHelloService
	SayHello() string

	//  discovery service from consul by serviceName
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}

var ErrNotServiceInstances = errors.New("instances are not existed")

type DiscoveryServiceImpl struct {
	discoveryClient discover.DiscoveryClient
}
// 构造函数，实现service接口
func NewDiscoveryServiceImpl(discoveryClient discover.DiscoveryClient) Service {
	return &DiscoveryServiceImpl{
		discoveryClient: discoveryClient,
	}
}

func (*DiscoveryServiceImpl) SayHello() string {
	return "Hello World!"
}

func (service *DiscoveryServiceImpl) DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error) {
	//从consul中根据服务名获取服务示例列表
	instances := service.discoveryClient.DiscoverServices(serviceName, config.Logger)

	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}

// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true
func (*DiscoveryServiceImpl) HealthCheck() bool {
	return true
}
