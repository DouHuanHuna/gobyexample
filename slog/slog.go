package main

import (
	"log/slog"
	"os"
)

func main() {
	// 创建一个新的 Logger 实例，输出到标准输出（终端）
	logger := slog.New(slog.NewTextHandler(os.Stdout))

	// 使用 Logger 输出不同级别的日志信息
	logger.Info("这是一个信息日志", "key", "value")
	logger.Warn("这是一个警告日志", "module", "example")
	logger.Error("这是一个错误日志", "code", 500)

	// 示例：日志记录带有结构化数据
	logger.Info("结构化日志", "user", "John Doe", "action", "login")
}
