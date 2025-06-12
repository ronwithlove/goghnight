package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/yourusername/goghnight/models"
)

// MessageService 消息服务
type MessageService struct{}

// GetLatestMessage 获取最新的消息
func (s *MessageService) GetLatestMessage() (*models.Message, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")

	// 构建请求URL
	url := fmt.Sprintf("%s/rest/v1/messages?select=*&order=created_at.desc&limit=1", supabaseURL)

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API请求失败: %s", string(body))
	}

	// 解析响应
	var messages []models.Message
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, err
	}

	// 如果没有消息，返回错误
	if len(messages) == 0 {
		return nil, fmt.Errorf("没有找到消息")
	}

	return &messages[0], nil
}

// GetAllMessages 获取所有消息
func (s *MessageService) GetAllMessages() ([]models.Message, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")

	// 构建请求URL
	url := fmt.Sprintf("%s/rest/v1/messages?select=*&order=created_at.desc", supabaseURL)

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API请求失败: %s", string(body))
	}

	// 解析响应
	var messages []models.Message
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// CreateMessage 创建新消息
func (s *MessageService) CreateMessage(content string) (*models.Message, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")

	// 构建请求URL
	url := fmt.Sprintf("%s/rest/v1/messages", supabaseURL)

	// 构建请求体
	message := models.Message{
		Content: content,
	}
	requestBody, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查响应状态
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("创建消息失败: %s", string(body))
	}

	// 解析响应
	var messages []models.Message
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, err
	}

	// 返回创建的消息
	if len(messages) > 0 {
		return &messages[0], nil
	}

	return nil, fmt.Errorf("创建消息失败：没有返回数据")
}
