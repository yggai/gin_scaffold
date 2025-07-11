package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/spf13/cobra"
)

// 测试创建命令基本功能
func TestCreateCommand(t *testing.T) {
	// 准备测试环境
	cmd := NewCreateCmd()
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	
	// 执行命令（不带参数，应该显示帮助）
	err := cmd.Execute()
	assert.NoError(t, err, "create命令不带参数应该显示帮助")
	
	output := b.String()
	assert.Contains(t, output, "create", "输出应该包含命令名称")
}

// 测试创建命令帮助
func TestCreateCommandHelp(t *testing.T) {
	cmd := NewCreateCmd()
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	
	err := cmd.Execute()
	assert.NoError(t, err, "执行帮助命令失败")
	
	output := b.String()
	assert.Contains(t, output, "创建", "帮助输出应该包含命令描述")
	assert.Contains(t, output, "用法:", "帮助输出应该包含用法信息")
	assert.Contains(t, output, "可用命令:", "帮助输出应该包含子命令列表")
}

// 测试控制器创建命令
func TestCreateControllerCommand(t *testing.T) {
	// 创建临时目录作为测试项目
	tempDir, err := os.MkdirTemp("", "gs-create-test-")
	require.NoError(t, err, "无法创建临时测试目录")
	defer os.RemoveAll(tempDir)
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "无法获取当前工作目录")
	defer os.Chdir(originalDir)
	
	err = os.Chdir(tempDir)
	require.NoError(t, err, "无法切换到临时目录")
	
	// 模拟一个简单的Go模块环境
	err = os.WriteFile("go.mod", []byte("module example.com/testapp"), 0644)
	require.NoError(t, err, "无法创建go.mod文件")
	
	// 获取模板目录
	templatesDirFromEnv := os.Getenv("GS_TEMPLATES_DIR")
	var templatesDir string
	if templatesDirFromEnv != "" {
		templatesDir = templatesDirFromEnv
	} else {
		// 尝试使用默认位置
		templatesDir = filepath.Join(originalDir, "templates")
	}
	
	// 如果模板目录不存在，则跳过实际创建测试
	if _, err := os.Stat(filepath.Join(templatesDir, "component", "controller")); os.IsNotExist(err) {
		t.Skip("跳过控制器创建测试：找不到模板目录")
	}
	
	// 创建子命令
	createCmd := NewCreateCmd()
	cmd := createCmd.Commands()[0] // 假设controller是第一个子命令
	
	// 如果无法找到controller子命令，则手动添加
	if cmd.Use != "controller" {
		for _, c := range createCmd.Commands() {
			if c.Use == "controller" {
				cmd = c
				break
			}
		}
	}
	
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"User"})
	
	// 设置模板目录环境变量
	os.Setenv("GS_TEMPLATES_DIR", templatesDir)
	defer os.Unsetenv("GS_TEMPLATES_DIR")
	
	// 执行命令
	err = cmd.Execute()
	if err != nil {
		t.Logf("命令输出: %s", b.String())
	}
	assert.NoError(t, err, "控制器创建命令失败")
	
	// 验证控制器文件是否已生成
	controllerFile := filepath.Join(tempDir, "controllers", "user_controller.go")
	_, err = os.Stat(controllerFile)
	assert.False(t, os.IsNotExist(err), "未生成控制器文件")
	
	// 验证输出
	output := b.String()
	assert.Contains(t, output, "成功", "输出应该表明控制器创建成功")
}

// 测试模型创建命令
func TestCreateModelCommand(t *testing.T) {
	// 创建临时目录作为测试项目
	tempDir, err := os.MkdirTemp("", "gs-create-test-")
	require.NoError(t, err, "无法创建临时测试目录")
	defer os.RemoveAll(tempDir)
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "无法获取当前工作目录")
	defer os.Chdir(originalDir)
	
	err = os.Chdir(tempDir)
	require.NoError(t, err, "无法切换到临时目录")
	
	// 模拟一个简单的Go模块环境
	err = os.WriteFile("go.mod", []byte("module example.com/testapp"), 0644)
	require.NoError(t, err, "无法创建go.mod文件")
	
	// 获取模板目录
	templatesDirFromEnv := os.Getenv("GS_TEMPLATES_DIR")
	var templatesDir string
	if templatesDirFromEnv != "" {
		templatesDir = templatesDirFromEnv
	} else {
		// 尝试使用默认位置
		templatesDir = filepath.Join(originalDir, "templates")
	}
	
	// 如果模板目录不存在，则跳过实际创建测试
	if _, err := os.Stat(filepath.Join(templatesDir, "component", "model")); os.IsNotExist(err) {
		t.Skip("跳过模型创建测试：找不到模板目录")
	}
	
	// 查找模型子命令
	createCmd := NewCreateCmd()
	var cmd *cobra.Command
	for _, c := range createCmd.Commands() {
		if c.Use == "model" {
			cmd = c
			break
		}
	}
	
	if cmd == nil {
		t.Skip("跳过模型创建测试：找不到model子命令")
	}
	
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"User"})
	
	// 设置模板目录环境变量
	os.Setenv("GS_TEMPLATES_DIR", templatesDir)
	defer os.Unsetenv("GS_TEMPLATES_DIR")
	
	// 执行命令
	err = cmd.Execute()
	if err != nil {
		t.Logf("命令输出: %s", b.String())
	}
	assert.NoError(t, err, "模型创建命令失败")
	
	// 验证模型文件是否已生成
	modelFile := filepath.Join(tempDir, "models", "user.go")
	_, err = os.Stat(modelFile)
	assert.False(t, os.IsNotExist(err), "未生成模型文件")
	
	// 验证输出
	output := b.String()
	assert.Contains(t, output, "成功", "输出应该表明模型创建成功")
} 