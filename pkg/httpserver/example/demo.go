package main

import (
	"fmt"
	"io"
	"net/http"

	hs "github.com/shengyanli1982/toolkit/pkg/httpserver"
)

// healthCheckFuncHandler 是一个处理器函数，它将 HTTP 状态码设置为 OK，并向响应写入 "lee"
// healthCheckFuncHandler is a handler function that sets the HTTP status code to OK and writes "lee" to the response
func healthCheckFuncHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("lee"))
}

func main() {
	// 创建一个新的 TinyHttpServer 实例，监听默认端口，没有路由，使用 healthCheckFuncHandler 作为处理器
	// Create a new TinyHttpServer instance, listen on the default port, no routes, use healthCheckFuncHandler as the handler
	srv := hs.NewTinyHttpServer(hs.DefaultListenPort, nil, healthCheckFuncHandler)

	// 使用默认的 HTTP 客户端发送 GET 请求到 "http://localhost:8080/ping"
	// Use the default HTTP client to send a GET request to "http://localhost:8080/ping"
	if httpResp, err := http.DefaultClient.Get("http://localhost:8080/ping"); err == nil {
		// 如果 HTTP 状态码是 OK
		// If the HTTP status code is OK
		if httpResp.StatusCode == http.StatusOK {
			// 读取响应体
			// Read the response body
			if content, err := io.ReadAll(httpResp.Body); err != nil {
				// 如果读取响应体失败，打印错误信息
				// If reading the response body fails, print the error message
				fmt.Printf("Read response body failed !!\n")
			} else {
				// 打印 HTTP 状态码和响应体
				// Print the HTTP status code and the response body
				fmt.Printf(">>>> %d:%s\n", httpResp.StatusCode, content)
			}
		} else {
			// 如果 HTTP 状态码不是 OK，打印错误信息
			// If the HTTP status code is not OK, print the error message
			fmt.Printf("Server is not running !!\n")
		}

		// 关闭响应体
		// Close the response body
		httpResp.Body.Close()
	}

	// 停止服务器
	// Stop the server
	srv.Stop()
}
