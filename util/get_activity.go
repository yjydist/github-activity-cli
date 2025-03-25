package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/yjydist/github-activity-cli/model"
)

// GetActivity 获取用户的活动
func GetActivity(username string) ([]model.Activity, error) {
	url := "https://api.github.com/users/" + username + "/events" // 构造请求 URL
	response, err := http.Get(url) // 发送 HTTP GET 请求
	
	// 检查错误
	if err != nil {
		log.Printf("failed to fetch data: %v", err)
		return nil, err
	} 

	// 检查 HTTP 响应状态码
	if response.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %d", response.StatusCode)
		return nil, fmt.Errorf("failed to fetch data: status code %d", response.StatusCode)
	}

	// 延迟关闭响应体
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(response.Body)

	

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return nil, err
	}

	var activities []model.Activity
	err = json.Unmarshal(responseData, &activities)
	if err != nil {
		log.Printf("failed to unmarshal response data: %v", err)
		return nil, err
	}

	return activities, nil
}
