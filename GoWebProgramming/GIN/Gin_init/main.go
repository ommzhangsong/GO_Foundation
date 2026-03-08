package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func index(context *gin.Context) {

	context.JSON(200, Response{Code: 0, Msg: "hello", Data: true})
}
func main() {
	//init
	gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//// ✅ 手动添加 Logger 和 Recovery 中间件
	//r.Use(gin.Logger())   // 打印请求日志
	//r.Use(gin.Recovery()) // 崩溃恢复
	r := gin.Default()
	//挂路由
	// ✅ 仅信任你的真实代理 IP（例如 Nginx 所在服务器 IP）
	// 单个 IP：
	_ = r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.100"})
	// 或信任一个网段（如本地局域网）：
	// r.TrustedProxies = []string{"127.0.0.1/8", "192.168.0.0/16"}
	r.GET("/index", index)
	//运行
	_ = r.Run(":8080")
}
