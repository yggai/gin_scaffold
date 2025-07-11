package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileExists 检查文件是否存在
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// EnsureDir 确保目录存在，如果不存在则创建
func EnsureDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, 0755)
	}
	return nil
}

// GetProjectRoot 获取当前项目的根目录
func GetProjectRoot() (string, error) {
	// 首先检查当前目录是否是项目根目录（是否存在go.mod文件）
	if FileExists("go.mod") {
		return ".", nil
	}

	// 向上递归查找go.mod文件
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("无法获取当前工作目录: %v", err)
	}

	for {
		// 检查当前目录是否有go.mod文件
		if FileExists(filepath.Join(currentDir, "go.mod")) {
			return currentDir, nil
		}

		// 获取父目录
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// 已经到达根目录，但仍然没有找到go.mod文件
			break
		}
		currentDir = parentDir
	}

	return "", fmt.Errorf("无法找到项目根目录（包含go.mod的目录）")
}

// SanitizeName 清理名称，确保符合Go的标识符规范
func SanitizeName(name string) string {
	// 移除非字母数字字符
	result := ""
	for _, char := range name {
		if (char >= 'a' && char <= 'z') || 
		   (char >= 'A' && char <= 'Z') || 
		   (char >= '0' && char <= '9') || 
		   char == '_' {
			result += string(char)
		} else if char == '-' || char == ' ' || char == '.' {
			result += "_"
		}
	}
	
	// 确保首字符不是数字
	if len(result) > 0 && result[0] >= '0' && result[0] <= '9' {
		result = "_" + result
	}
	
	return result
}

// FormatPackageName 格式化包名（确保全小写，无特殊字符）
func FormatPackageName(name string) string {
	return strings.ToLower(SanitizeName(name))
} 