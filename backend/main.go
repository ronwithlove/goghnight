package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	// 创建gin路由器
	r := gin.Default()

	// API路由
	api := r.Group("/api")
	{
		api.GET("/message", getMessage)
	}

	// 设置CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// 用CORS handler包裹gin
	handler := c.Handler(r)

	// 启动服务器
	log.Println("后端服务器启动在端口 8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}

// getMessage 返回消息的API端点
func getMessage(c *gin.Context) {
	message := gin.H{
		"message": "Hello from Go Backend! 🚀",
		"status":  "success",
	}
	c.JSON(http.StatusOK, message)
}
