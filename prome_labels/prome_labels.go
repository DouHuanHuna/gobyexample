package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	// 创建一个新的 CounterVec 指标，记录 HTTP 请求的总数，标签包括 method 和 status
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "status"}, // 指定标签
	)

	// 注册指标
	prometheus.MustRegister(httpRequestsTotal)

	// 模拟 HTTP 请求的处理
	go func() {
		// 模拟 3 个 HTTP 请求，分别使用不同的 HTTP 方法和状态码
		httpRequestsTotal.WithLabelValues("GET", "200").Inc()  // GET 请求成功
		httpRequestsTotal.WithLabelValues("GET", "404").Inc()  // GET 请求失败
		httpRequestsTotal.WithLabelValues("POST", "200").Inc() // POST 请求成功
		httpRequestsTotal.WithLabelValues("POST", "500").Inc() // POST 请求失败
		httpRequestsTotal.WithLabelValues("POST", "500").Inc() // POST 请求失败
		httpRequestsTotal.WithLabelValues("POST", "500").Inc() // POST 请求失败
	}()

	// 启动 HTTP 服务器暴露 Prometheus 指标
	http.Handle("/metrics", promhttp.Handler()) // 这个路由会暴露 Prometheus 指标
	log.Fatal(http.ListenAndServe(":8089", nil))
}
