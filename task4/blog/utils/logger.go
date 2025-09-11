package utils

//这段代码是整个项目的「日志基础设施」，只做两件极小的事：
//初始化一个 zap 生产级日志器
//把它包装成 全局变量 L，供其他包直接调用

import "go.uber.org/zap"

var L *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	L = logger.Sugar()
}
