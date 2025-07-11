package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试根命令执行
func TestRootCommand(t *testing.T) {
	// 准备测试环境
	cmd := NewRootCmd()
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetErr(b)
	
	// 执行命令
	err := cmd.Execute()
	assert.NoError(t, err, "根命令执行失败")
	
	// 验证输出包含有关gs工具的信息
	output := b.String()
	assert.Contains(t, output, "gs", "输出应该包含工具名称")
}

// 测试命令行标志
func TestRootCommandFlags(t *testing.T) {
	cmd := NewRootCmd()
	
	// 验证版本标志的存在
	versionFlag := cmd.Flag("version")
	assert.NotNil(t, versionFlag, "版本标志应该存在")
	
	// 测试版本标志
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--version"})
	
	err := cmd.Execute()
	assert.NoError(t, err, "执行带有版本标志的命令失败")
	
	output := b.String()
	assert.Contains(t, output, "gs版本", "输出应该包含版本信息")
}

// 测试获取命令帮助
func TestRootCommandHelp(t *testing.T) {
	cmd := NewRootCmd()
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	
	err := cmd.Execute()
	assert.NoError(t, err, "执行帮助命令失败")
	
	output := b.String()
	assert.Contains(t, output, "用法:", "帮助输出应该包含用法信息")
	assert.Contains(t, output, "可用命令:", "帮助输出应该包含可用命令列表")
	assert.Contains(t, output, "标志:", "帮助输出应该包含标志信息")
} 