package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试初始化命令基本功能
func TestInitCommand(t *testing.T) {
	// 准备测试环境
	cmd := NewInitCmd()
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	
	// 执行命令（不带参数，应该显示帮助或错误）
	err := cmd.Execute()
	assert.Error(t, err, "在没有提供项目名称时，init命令应该失败")
	
	output := b.String()
	assert.Contains(t, output, "错误", "输出应该包含错误信息")
}

// 测试初始化命令帮助
func TestInitCommandHelp(t *testing.T) {
	cmd := NewInitCmd()
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	
	err := cmd.Execute()
	assert.NoError(t, err, "执行帮助命令失败")
	
	output := b.String()
	assert.Contains(t, output, "初始化", "帮助输出应该包含命令描述")
	assert.Contains(t, output, "用法:", "帮助输出应该包含用法信息")
}

// 测试实际项目初始化
func TestInitCommandWithName(t *testing.T) {
	// 创建临时目录作为测试工作区
	tempDir, err := os.MkdirTemp("", "gs-init-test-")
	require.NoError(t, err, "无法创建临时测试目录")
	defer os.RemoveAll(tempDir)
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "无法获取当前工作目录")
	defer os.Chdir(originalDir)
	
	err = os.Chdir(tempDir)
	require.NoError(t, err, "无法切换到临时目录")
	
	// 获取模板目录
	templatesDirFromEnv := os.Getenv("GS_TEMPLATES_DIR")
	var templatesDir string
	if templatesDirFromEnv != "" {
		templatesDir = templatesDirFromEnv
	} else {
		// 尝试使用默认位置
		templatesDir = filepath.Join(originalDir, "templates")
	}
	
	// 如果模板目录不存在，则跳过实际初始化测试
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		t.Skip("跳过实际初始化测试：找不到模板目录")
	}
	
	// 准备命令
	cmd := NewInitCmd()
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"test-app", "--module", "github.com/username/test-app"})
	
	// 执行命令
	err = cmd.Execute()
	assert.NoError(t, err, "项目初始化失败")
	
	// 验证关键文件是否已生成
	expectedFiles := []string{
		"main.go",
		"go.mod",
		"README.md",
	}
	
	for _, file := range expectedFiles {
		filePath := filepath.Join(tempDir, "test-app", file)
		_, err := os.Stat(filePath)
		assert.False(t, os.IsNotExist(err), "未生成文件: %s", file)
	}
	
	// 验证输出
	output := b.String()
	assert.Contains(t, output, "成功", "输出应该表明初始化成功")
} 