package service

import (
	"fmt"
	"strings"
	"xpass/tools"
)

func CheckEnvironment() bool {
	// 检查java版本（是否安装）
	jrever := tools.Run("java", "-version")
	if !strings.Contains(jrever, "1.8") {
		tools.Err("Environment Check failed! Cannot find jre 1.8 path, check your PATH or install new jre.", nil)
		return false
	}
	// 检查nacos

	// 检查redis

	// 检查端口是否被占用
	ports := []int{8002, 8004, 8006, 8008, 8010, 8012}

	for _, port := range ports {
		if !tools.PortCheck(port) {
			tools.Err(fmt.Sprintf("Port check failed, %d is used!", port), nil)
			return false
		}
	}

	return true
}
