// Package utils 一般响应在utils/response.go，目录下
package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(200, Response{
		Code: 500,
		Msg:  msg,
		Data: gin.H{},
	})
}
