package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpts "github.com/consul-discover/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoints endpts.DiscoveryEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter() //路由
	//定义处理器
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	//say-hello 接口
	r.Methods("GET").Path("/say-hello").Handler(kithttp.NewServer(
		endpoints.SayHelloEndpoint,
		decodeSayHelloRequest,
		encodeJsonResponse,
		options...,
	))
	//服务发现接口
	r.Methods("GET").Path("/discovery").Handler(kithttp.NewServer(
		endpoints.DiscoveryEndpoint,
		decodeDiscoveryRequest,
		encodeJsonResponse,
		options...,
	))

	// create health check handler 健康检查
	r.Methods("GET").Path("/health").Handler(kithttp.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeJsonResponse,
		options...,
	))

	return r
}

/*
decode 将http请求转化为可以接受的结构体
encode 将http返回的结构体转化为对应的http响应

 */

// decodeSayHelloRequest decode request params to struct
func decodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpts.SayHelloRequest{}, nil
}

// decodeDiscoveryRequest decode request params to struct
func decodeDiscoveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	serviceName := r.URL.Query().Get("serviceName")
	if serviceName == "" {
		return nil, ErrorBadRequest
	}
	return endpts.DiscoveryRequest{
		ServiceName: serviceName,
	}, nil
}

// decodeHealthCheckRequest decode request
func decodeHealthCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpts.HealthRequest{}, nil
}

// encodeJsonResponse encode response to return
func encodeJsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
