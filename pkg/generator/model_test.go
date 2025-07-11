package generator

import (
	"os"
	"path/filepath"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

// 测试模型数据结构
func TestModelData(t *testing.T) {
	data := ModelData{
		Name:      "User",
		TableName: "users",
		VarName:   "user",
		Package:   "myapp",
	}
	
	assert.Equal(t, "User", data.Name, "Name字段设置不正确")
	assert.Equal(t, "users", data.TableName, "TableName字段设置不正确")
	assert.Equal(t, "user", data.VarName, "VarName字段设置不正确")
	assert.Equal(t, "myapp", data.Package, "Package字段设置不正确")
}

// 测试生成模型
func TestGenerateModel(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("无法获取当前工作目录: %v", err)
	}
	defer os.Chdir(originalDir)
	
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("无法切换到临时目录: %v", err)
	}
	
	// 创建templates目录结构
	templatesDir := filepath.Join(tempDir, "templates", "component", "model")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("无法创建模板目录: %v", err)
	}
	
	// 创建测试模板
	templateContent := `package models

import (
	"time"
)

// {{.Name}} 表示{{.Name}}模型
type {{.Name}} struct {
	ID        uint      ` + "`json:\"id\" gorm:\"primaryKey\"`" + `
	Name      string    ` + "`json:\"name\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
}

// TableName 指定表名
func ({{.Name}}) TableName() string {
	return "{{.TableName}}"
}
`
	templatePath := filepath.Join(templatesDir, "model.go.tmpl")
	if err := os.WriteFile(templatePath, []byte(templateContent), 0644); err != nil {
		t.Fatalf("无法创建测试模板文件: %v", err)
	}
	
	// 创建生成器
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试生成模型
	err = g.GenerateModel("User", "myapp")
	if err != nil {
		t.Fatalf("生成模型失败: %v", err)
	}
	
	// 验证模型文件是否已生成
	modelPath := filepath.Join(tempDir, "models", "user.go")
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		t.Fatalf("模型文件未生成: %s", modelPath)
	}
	
	// 验证模型内容
	content, err := os.ReadFile(modelPath)
	if err != nil {
		t.Fatalf("无法读取生成的模型文件: %v", err)
	}
	
	// 检查内容是否符合预期
	expectedContent := `package models

import (
	"time"
)

// User 表示User模型
type User struct {
	ID        uint      ` + "`json:\"id\" gorm:\"primaryKey\"`" + `
	Name      string    ` + "`json:\"name\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
`
	if string(content) != expectedContent {
		t.Errorf("生成的模型内容不符合预期\n期望:\n%s\n实际:\n%s", expectedContent, string(content))
	}
}

// 测试生成模型 - 错误情况
func TestGenerateModel_Errors(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 创建生成器，但不创建模板目录
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试模板不存在的情况
	err := g.GenerateModel("User", "myapp")
	if err == nil {
		t.Error("期望在模板不存在时返回错误，但没有")
	}
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("无法获取当前工作目录: %v", err)
	}
	defer os.Chdir(originalDir)
	
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("无法切换到临时目录: %v", err)
	}
	
	// 创建模板目录
	templatesDir := filepath.Join(tempDir, "templates", "component", "model")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("无法创建模板目录: %v", err)
	}
	
	// 创建测试模板
	templatePath := filepath.Join(templatesDir, "model.go.tmpl")
	if err := os.WriteFile(templatePath, []byte("template content"), 0644); err != nil {
		t.Fatalf("无法创建测试模板文件: %v", err)
	}
	
	// 创建models目录和已存在的模型文件
	modelsDir := filepath.Join(tempDir, "models")
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		t.Fatalf("无法创建models目录: %v", err)
	}
	
	// 创建已存在的模型文件
	modelPath := filepath.Join(modelsDir, "user.go")
	if err := os.WriteFile(modelPath, []byte("already exists"), 0644); err != nil {
		t.Fatalf("无法创建已存在的模型文件: %v", err)
	}
	
	// 测试模型文件已存在的情况
	err = g.GenerateModel("User", "myapp")
	if err == nil {
		t.Error("期望在模型文件已存在时返回错误，但没有")
	}
} 