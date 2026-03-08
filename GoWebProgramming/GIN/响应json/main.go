package main

import "github.com/gin-gonic/gin"
import "go_foundation/GoWebProgramming/GIN/utils"

func index(c *gin.Context) {
	//data := gin.H{
	//	"code": 0,
	//	"msg":  "success",
	//	"data": gin.H{
	//		"name": "张三",
	//		"age":  20,
	//	},
	//}
	//context.JSON(200, data)
	user := gin.H{"name": "张三",
		"age": 20}
	utils.Success(c, user)
	c.String(200, "hello")

}
func main() {

	r := gin.Default()
	r.GET("/index", index)
	r.Run("localhost:8080")
}
