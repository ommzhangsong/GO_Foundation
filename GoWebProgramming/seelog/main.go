package main

import (
	"github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

func init() {
	// 配置：输出到滚动文件 + 控制台，级别为 Debug
	logger, _ = seelog.LoggerFromConfigAsString(`
		<seelog minlevel="debug">
			<outputs formatid="main">
				<!-- 滚动文件配置：按大小切割（100MB/文件），保留10个备份 -->
				<rollingfile filename="app.log" maxsize="104857600" maxrolls="10" />
				<console /> <!-- 同时输出到控制台 -->
			</outputs>
			<formats>
				<format id="main" format="%Date(%Y-%m-%d) %Time(%H:%M:%S) [%Level] [%Func] %Msg%n" />
			</formats>
		</seelog>
	`)
}

func cleanup() {
	logger.Flush()
	logger.Close()
}

func main() {
	defer cleanup()

	// 输出 Debug 级别日志（minlevel=debug，会输出）
	logger.Debug("调试信息：请求参数为 {\"user_id\":\"U001\"}")
	logger.Info("业务日志：创建订单成功，订单ID=N001")
	logger.Error("错误日志：调用第三方支付接口失败，err=超时")
}
