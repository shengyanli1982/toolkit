# Tiny Http Server

This is a tiny http server that can serve Kubernetes Pod. It is written in Go and uses the standard library. Just only support health check and metrics.

## Installation

```bash
go get github.com/shengyanli1982/toolkit/pkg/httpserver
```

## Quick Start

`httpserver` just needs to be imported and used directly.

**Example**

```go
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

	fmt.Println("==================== [Split Line] ====================")

	// 使用默认的 HTTP 客户端发送 GET 请求到 "http://localhost:8080/metrics"
	// Use the default HTTP client to send a GET request to "http://localhost:8080/metrics"
	if httpResp, err := http.DefaultClient.Get("http://localhost:8080/metrics"); err == nil {
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
```

**Result**

```bash
$ go run demo.go
>>>> 200:lee
==================== [Split Line] ====================
>>>> 200:# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.21.13"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 137520
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 137520
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4251
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 2.452944e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 137520
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 2.220032e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.548288e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 300
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 2.220032e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 3.76832e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 300
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 14400
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 46872
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 48888
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 991677
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 425984
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 425984
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.707664e+06
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 8
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
