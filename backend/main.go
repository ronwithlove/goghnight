package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/yourusername/goghnight/models"
	"github.com/yourusername/goghnight/services"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 文件，使用系统环境变量")
	}

	// 检查必要的环境变量
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("SUPABASE_URL 和 SUPABASE_ANON_KEY 环境变量必须设置")
	}

	log.Println("Supabase 配置加载成功")

	// 创建gin路由器
	r := gin.Default()

	// 创建服务实例
	messageService := &services.MessageService{}

	// API路由
	api := r.Group("/api")
	{
		api.GET("/message", getMessage(messageService))
		api.GET("/messages", getAllMessages(messageService))
		api.POST("/messages", createMessage(messageService))
	}

	// 获取允许的前端域名
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000" // 默认值
	}

	// 设置CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			frontendURL,
			"https://*.vercel.app", // 允许所有vercel.app子域名
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// 用CORS handler包裹gin
	handler := c.Handler(r)

	// 获取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	log.Printf("后端服务器启动在端口 %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}

// getMessage 获取最新消息的API端点
func getMessage(service *services.MessageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		message, err := service.GetLatestMessage()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MessageResponse{
				Message: "获取消息失败: " + err.Error(),
				Status:  "error",
			})
			return
		}

		c.JSON(http.StatusOK, models.MessageResponse{
			Message: message.Content,
			Status:  "success",
			Data:    message,
		})
	}
}

// getAllMessages 获取所有消息的API端点
func getAllMessages(service *services.MessageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		messages, err := service.GetAllMessages()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "获取消息列表失败: " + err.Error(),
				"status":  "error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "获取消息列表成功",
			"status":  "success",
			"data":    messages,
		})
	}
}

// createMessage 创建新消息的API端点
func createMessage(service *services.MessageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Content string `json:"content" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求参数错误",
				"status":  "error",
			})
			return
		}

		message, err := service.CreateMessage(request.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "创建消息失败: " + err.Error(),
				"status":  "error",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "创建消息成功",
			"status":  "success",
			"data":    message,
		})
	}
}
