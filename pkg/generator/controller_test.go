package generator

import (
	"os"
	"path/filepath"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试控制器数据结构
func TestControllerData(t *testing.T) {
	data := ControllerData{
		Name:         "User",
		PluralName:   "Users",
		ResourceName: "users",
		VarName:      "user",
		Package:      "myapp",
	}
	
	assert.Equal(t, "User", data.Name, "Name字段设置不正确")
	assert.Equal(t, "Users", data.PluralName, "PluralName字段设置不正确")
	assert.Equal(t, "users", data.ResourceName, "ResourceName字段设置不正确")
	assert.Equal(t, "user", data.VarName, "VarName字段设置不正确")
	assert.Equal(t, "myapp", data.Package, "Package字段设置不正确")
}

// 测试生成控制器
func TestGenerateController(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "无法获取当前工作目录")
	defer os.Chdir(originalDir)
	
	err = os.Chdir(tempDir)
	require.NoError(t, err, "无法切换到临时目录")
	
	// 创建templates目录结构
	templatesDir := filepath.Join(tempDir, "templates", "component", "controller")
	err = os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err, "无法创建模板目录")
	
	// 创建测试模板
	templateContent := `package controllers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// {{.Name}}Controller 处理{{.Name}}相关的HTTP请求
type {{.Name}}Controller struct {}

// Get{{.PluralName}} 获取所有{{.PluralName}}
func (c *{{.Name}}Controller) Get{{.PluralName}}(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "获取所有{{.PluralName}}"})
}
`
	templatePath := filepath.Join(templatesDir, "controller.go.tmpl")
	err = os.WriteFile(templatePath, []byte(templateContent), 0644)
	require.NoError(t, err, "无法创建测试模板文件")
	
	// 创建生成器
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试生成控制器
	err = g.GenerateController("User", "myapp")
	require.NoError(t, err, "生成控制器失败")
	
	// 验证控制器文件是否已生成
	controllerPath := filepath.Join(tempDir, "controllers", "user_controller.go")
	_, err = os.Stat(controllerPath)
	require.False(t, os.IsNotExist(err), "控制器文件未生成")
	
	// 验证控制器内容
	content, err := os.ReadFile(controllerPath)
	require.NoError(t, err, "无法读取生成的控制器文件")
	
	// 检查内容是否符合预期
	expectedContent := `package controllers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// UserController 处理User相关的HTTP请求
type UserController struct {}

// GetUsers 获取所有Users
func (c *UserController) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "获取所有Users"})
}
`
	assert.Equal(t, expectedContent, string(content), "生成的控制器内容不符合预期")
}

// 测试生成控制器 - 错误情况
func TestGenerateController_Errors(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 创建生成器，但不创建模板目录
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试模板不存在的情况
	err := g.GenerateController("User", "myapp")
	assert.Error(t, err, "期望在模板不存在时返回错误，但没有")
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "无法获取当前工作目录")
	defer os.Chdir(originalDir)
	
	err = os.Chdir(tempDir)
	require.NoError(t, err, "无法切换到临时目录")
	
	// 创建模板目录，但使控制器目录只读
	templatesDir := filepath.Join(tempDir, "templates", "component", "controller")
	err = os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err, "无法创建模板目录")
	
	// 创建测试模板
	templatePath := filepath.Join(templatesDir, "controller.go.tmpl")
	err = os.WriteFile(templatePath, []byte("template content"), 0644)
	require.NoError(t, err, "无法创建测试模板文件")
	
	// 创建controllers目录并设置为只读
	controllersDir := filepath.Join(tempDir, "controllers")
	err = os.MkdirAll(controllersDir, 0444)
	require.NoError(t, err, "无法创建controllers目录")
	
	// 在Windows上设置为只读不足以阻止写入
	// 这里我们测试文件已存在的情况
	controllerPath := filepath.Join(controllersDir, "user_controller.go")
	err = os.WriteFile(controllerPath, []byte("already exists"), 0644)
	require.NoError(t, err, "无法创建已存在的控制器文件")
	
	// 测试控制器文件已存在的情况
	err = g.GenerateController("User", "myapp")
	assert.Error(t, err, "期望在控制器文件已存在时返回错误，但没有")
} 