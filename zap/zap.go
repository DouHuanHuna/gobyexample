package main

import (
	"fmt"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 使用 logger 来记录不同级别的日志
	logger.Info("这是一个信息级别的日志",
		zap.String("module", "example"),
		zap.Int("attempt", 3),
	)

	logger.Warn("这是一个警告级别的日志",
		zap.String("module", "example"),
	)

	logger.Error("发生了一个错误",
		zap.String("module", "example"),
		zap.Error(fmt.Errorf("这是一个示例错误")),
	)
}
