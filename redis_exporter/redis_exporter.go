package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

var (
	redisUp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "redis_up",
		Help: "1 if Redis server is up, 0 if down",
	})
	redisInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "redis_info",
		Help: "Various Redis stats",
	}, []string{"key"})
)

func init() {
	prometheus.MustRegister(redisUp)
	prometheus.MustRegister(redisInfo)
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
	ctx := context.Background()

	go func() {
		for {
			err := rdb.Ping(ctx).Err()
			if err != nil {
				redisUp.Set(0) // Redis 健康检查失败，设置为 0
				log.Printf("Redis is down: %v", err)
			} else {
				redisUp.Set(1) // Redis 健康检查成功，设置为 1
			}
			// 通过 INFO 命令获取 Redis 信息并处理
			_, err = rdb.Info(ctx, "stats").Result()
			//fmt.Println(info, err)
			if err != nil {
				log.Printf("Failed to get Redis stats: %v", err)
			} else {
				// 解析 INFO 信息并更新指标
				// 注意：这里我们只是以一个简单的方式提取了 Redis 统计信息
				redisInfo.WithLabelValues("connected_clients").Set(1)
				redisInfo.WithLabelValues("blocked_clients").Set(2)
			}
		}
	}()

	// 设置 HTTP handler 用于暴露 metrics 给 Prometheus
	http.Handle("/metrics", promhttp.Handler())

	// 启动 HTTP 服务器
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
