package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试服务数据结构
func TestServiceData(t *testing.T) {
	data := ServiceData{
		Name:      "User",
		VarName:   "user",
		Package:   "myapp",
	}
	
	assert.Equal(t, "User", data.Name, "Name字段设置不正确")
	assert.Equal(t, "user", data.VarName, "VarName字段设置不正确")
	assert.Equal(t, "myapp", data.Package, "Package字段设置不正确")
}

// 测试生成服务
func TestGenerateService(t *testing.T) {
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
	templatesDir := filepath.Join(tempDir, "templates", "component", "service")
	err = os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err, "无法创建模板目录")
	
	// 创建测试模板
	templateContent := `package services

import (
	"{{.Package}}/models"
	"errors"
)

// {{.Name}}Service 处理{{.Name}}相关的业务逻辑
type {{.Name}}Service struct {}

// Get{{.Name}}ByID 根据ID获取{{.Name}}
func (s *{{.Name}}Service) Get{{.Name}}ByID(id uint) (*models.{{.Name}}, error) {
	// 模拟从数据库获取{{.Name}}
	if id == 0 {
		return nil, errors.New("ID不能为0")
	}
	
	return &models.{{.Name}}{
		ID: id,
		Name: "Test {{.Name}}",
	}, nil
}
`
	templatePath := filepath.Join(templatesDir, "service.go.tmpl")
	err = os.WriteFile(templatePath, []byte(templateContent), 0644)
	require.NoError(t, err, "无法创建测试模板文件")
	
	// 创建生成器
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试生成服务
	err = g.GenerateService("User", "myapp")
	require.NoError(t, err, "生成服务失败")
	
	// 验证服务文件是否已生成
	servicePath := filepath.Join(tempDir, "services", "user_service.go")
	_, err = os.Stat(servicePath)
	require.False(t, os.IsNotExist(err), "服务文件未生成")
	
	// 验证服务内容
	content, err := os.ReadFile(servicePath)
	require.NoError(t, err, "无法读取生成的服务文件")
	
	// 检查内容是否符合预期
	expectedContent := `package services

import (
	"myapp/models"
	"errors"
)

// UserService 处理User相关的业务逻辑
type UserService struct {}

// GetUserByID 根据ID获取User
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	// 模拟从数据库获取User
	if id == 0 {
		return nil, errors.New("ID不能为0")
	}
	
	return &models.User{
		ID: id,
		Name: "Test User",
	}, nil
}
`
	assert.Equal(t, expectedContent, string(content), "生成的服务内容不符合预期")
}

// 测试生成服务 - 错误情况
func TestGenerateService_Errors(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 创建生成器，但不创建模板目录
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试模板不存在的情况
	err := g.GenerateService("User", "myapp")
	assert.Error(t, err, "期望在模板不存在时返回错误，但没有")
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "无法获取当前工作目录")
	defer os.Chdir(originalDir)
	
	err = os.Chdir(tempDir)
	require.NoError(t, err, "无法切换到临时目录")
	
	// 创建模板目录
	templatesDir := filepath.Join(tempDir, "templates", "component", "service")
	err = os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err, "无法创建模板目录")
	
	// 创建测试模板
	templatePath := filepath.Join(templatesDir, "service.go.tmpl")
	err = os.WriteFile(templatePath, []byte("template content"), 0644)
	require.NoError(t, err, "无法创建测试模板文件")
	
	// 创建services目录和已存在的服务文件
	servicesDir := filepath.Join(tempDir, "services")
	err = os.MkdirAll(servicesDir, 0755)
	require.NoError(t, err, "无法创建services目录")
	
	// 创建已存在的服务文件
	servicePath := filepath.Join(servicesDir, "user_service.go")
	err = os.WriteFile(servicePath, []byte("already exists"), 0644)
	require.NoError(t, err, "无法创建已存在的服务文件")
	
	// 测试服务文件已存在的情况
	err = g.GenerateService("User", "myapp")
	assert.Error(t, err, "期望在服务文件已存在时返回错误，但没有")
} 