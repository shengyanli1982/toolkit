package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// defaultHealthCheckUrl 是默认的健康检查URL
	// defaultHealthCheckUrl is the default health check URL
	defaultHealthCheckUrl = "/ping"

	// defaultMetricsUrl 是默认的 metrics URL
	// defaultMetricsUrl is the default metrics URL
	defaultMetricsUrl = "/metrics"

	// listenAddr 是服务器监听的地址
	// listenAddr is the address the server listens on
	listenAddr = "0.0.0.0"

	// DefaultListenPort 是服务器监听的默认端口
	// DefaultListenPort is the default port the server listens on
	DefaultListenPort = 8080

	// defaultIdleTimeout 是默认的空闲超时时间（秒）
	// defaultIdleTimeout is the default idle timeout in seconds
	defaultIdleTimeout = 120
)

// ok 是默认的健康检查响应
// ok is the default health check response
var ok = []byte("ok!!")

// defaultHealthCheckFuncHandler 是默认的健康检查处理函数
// defaultHealthCheckFuncHandler is the default health check handler function
func defaultHealthCheckFuncHandler(w http.ResponseWriter, _ *http.Request) {
	// 设置响应状态码为200
	// Set the response status code to 200
	w.WriteHeader(http.StatusOK)

	// 写入健康检查响应
	// Write the health check response
	_, _ = w.Write(ok)
}

// TinyHttpServer 定义了一个简单的HTTP服务器
// TinyHttpServer defines a simple HTTP server
type TinyHttpServer struct {
	// port 是服务器监听的端口
	// port is the port the server listens on
	port uint16

	// httpsvr 是http服务器实例
	// httpsvr is the http server instance
	httpsvr *http.Server

	// once 用于确保服务器只启动一次
	// once is used to ensure the server only starts once
	once *sync.Once

	// wg 用于等待服务器关闭
	// wg is used to wait for the server to shut down
	wg *sync.WaitGroup

	// log 是服务器的日志记录器
	// log is the server's logger
	log Logger
}

// NewTinyHttpServer 创建一个新的 TinyHttpServer 实例
// NewTinyHttpServer creates a new TinyHttpServer instance
func NewTinyHttpServer(port uint16, logger Logger, hcFunc func(w http.ResponseWriter, r *http.Request)) *TinyHttpServer {
	// 创建一个新的 HTTP 服务路由
	// Create a new HTTP service route
	mux := http.NewServeMux()

	// 如果 hcFunc 不为空，则将其作为默认的健康检查 URL 的处理函数
	// If hcFunc is not nil, use it as the handler function for the default health check URL
	if hcFunc != nil {
		mux.HandleFunc(defaultHealthCheckUrl, hcFunc)
	} else {
		// 否则，使用默认的健康检查函数处理器
		// Otherwise, use the default health check function handler
		mux.HandleFunc(defaultHealthCheckUrl, defaultHealthCheckFuncHandler)
	}

	// 添加默认的 metrics URL
	// Add the default metrics URL
	mux.HandleFunc(defaultMetricsUrl, promhttp.Handler().ServeHTTP)

	// 如果 logger 为空，则使用默认的 logger
	// If logger is nil, use the default logger
	if logger == nil {
		logger = &defaultLogger{}
	}

	// 创建一个新的 TinyHttpServer 实例
	// Create a new TinyHttpServer instance
	srv := &TinyHttpServer{
		// 端口号
		// Port number
		port: port,

		// sync.Once 用于只执行一次初始化
		// sync.Once is used for one-time initialization
		once: &sync.Once{},

		// sync.WaitGroup 用于等待所有的 goroutine 完成
		// sync.WaitGroup is used to wait for all goroutines to complete
		wg: &sync.WaitGroup{},

		// 日志记录器
		// Logger
		log: logger,

		// http.Server 是一个 HTTP 服务器
		// http.Server is an HTTP server
		httpsvr: &http.Server{
			// 服务器的监听地址
			// The listening address of the server
			Addr: fmt.Sprintf("%s:%d", listenAddr, port),

			// 处理 HTTP 请求的路由器
			// Router that handles HTTP requests
			Handler: mux,

			// 服务器空闲超时时间
			// Server idle timeout
			IdleTimeout: time.Second * defaultIdleTimeout,

			// 读取请求体的超时时间
			// Timeout for reading the request body
			ReadTimeout: time.Second * defaultIdleTimeout,

			// 读取请求头的超时时间
			// Timeout for reading the request header
			ReadHeaderTimeout: time.Duration(defaultIdleTimeout/2) * time.Second,

			// 写入响应的超时时间
			// Timeout for writing the response
			WriteTimeout: time.Second * defaultIdleTimeout,
		},
	}

	// 增加等待组的计数
	// Increase the count of the wait group
	srv.wg.Add(1)
	// 在一个新的 goroutine 中启动 HTTP 服务器
	// Start the HTTP server in a new goroutine
	go func() {
		defer srv.wg.Done()
		if err := srv.httpsvr.ListenAndServe(); err != nil {
			// 如果 HTTP 服务器启动失败，打印错误信息
			// If the HTTP server fails to start, print the error message
			srv.log.Errorf("http server start failed: %v\n", err)
		}
	}()

	// 返回新创建的 TinyHttpServer 实例
	// Return the newly created TinyHttpServer instance
	return srv
}

// Stop 停止服务器
// Stop stops the server
func (s *TinyHttpServer) Stop() {
	// 使用 sync.Once 确保服务器只停止一次
	// Use sync.Once to ensure the server only stops once
	s.once.Do(func() {
		// 创建一个5秒的超时上下文
		// Create a timeout context of 5 seconds
		ctx, cannel := context.WithTimeout(context.Background(), time.Second*5)

		// 尝试关闭http服务器
		// Try to shutdown the http server
		if err := s.httpsvr.Shutdown(ctx); err != nil {
			// 如果关闭失败，记录错误日志
			// If shutdown fails, log the error
			s.log.Errorf("http server stop failed` %v\n", err)
		}

		// 取消上下文
		// Cancel the context
		cannel()
	})
}
