package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type customCollector struct {
	desc *prometheus.Desc
}

func (c *customCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *customCollector) Collect(ch chan<- prometheus.Metric) {
	value := 42.0
	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.GaugeValue,
		value,
		"example_label_value",
	)
}

func main() {
	// 注销默认收集器
	prometheus.Unregister(collectors.NewGoCollector())
	//prometheus.Unregister(collectors.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	// 注册自定义收集器
	desc := prometheus.NewDesc(
		"example_custom_metric",
		"A custom metric example",
		[]string{"example_label"},
		nil,
	)
	collector := &customCollector{desc: desc}
	prometheus.MustRegister(collector)

	// 启动 HTTP 服务
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
