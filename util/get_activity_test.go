package util_test

import (
	"fmt"
	"testing"

	"github.com/yjydist/github-activity-cli/util"
)

func TestGetActivity(t *testing.T) {
	// 模拟输入
	username := "yjydist"

	// 调用被测试函数
	activities, err:= util.GetActivity(username)

	// 检查错误是否为 nil
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 检查返回值是否符合预期
	if len(activities) == 0 {
		t.Errorf("expected activities, got none")
	}

	fmt.Println(activities)
}
