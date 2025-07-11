package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试项目数据结构
func TestProjectData(t *testing.T) {
	data := ProjectData{
		Name:    "test-app",
		Module:  "github.com/username/test-app",
		Version: "v0.1.0",
	}
	
	assert.Equal(t, "test-app", data.Name, "Name字段设置不正确")
	assert.Equal(t, "github.com/username/test-app", data.Module, "Module字段设置不正确")
	assert.Equal(t, "v0.1.0", data.Version, "Version字段设置不正确")
}

// 测试初始化项目
func TestInitProject(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 切换到临时目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "无法获取当前工作目录")
	defer os.Chdir(originalDir)
	
	err = os.Chdir(tempDir)
	require.NoError(t, err, "无法切换到临时目录")
	
	// 创建templates目录结构和项目模板
	projectTemplatesDir := filepath.Join(tempDir, "templates", "project")
	err = os.MkdirAll(projectTemplatesDir, 0755)
	require.NoError(t, err, "无法创建项目模板目录")
	
	// 创建主要项目文件模板
	err = createProjectTemplates(t, projectTemplatesDir)
	require.NoError(t, err, "无法创建项目模板文件")
	
	// 创建生成器
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 创建项目目录
	projectDir := filepath.Join(tempDir, "my-app")
	err = os.Mkdir(projectDir, 0755)
	require.NoError(t, err, "无法创建项目目录")
	
	// 测试初始化项目
	err = os.Chdir(projectDir)
	require.NoError(t, err, "无法切换到项目目录")
	
	err = g.InitProject("my-app", "github.com/username/my-app")
	require.NoError(t, err, "项目初始化失败")
	
	// 验证主要文件是否已生成
	filesToCheck := []string{
		"main.go",
		"go.mod",
		"README.md",
		"config/config.go",
		"controllers/.gitkeep",
		"models/.gitkeep",
		"services/.gitkeep",
		"routers/router.go",
	}
	
	for _, file := range filesToCheck {
		_, err := os.Stat(filepath.Join(projectDir, file))
		assert.False(t, os.IsNotExist(err), "项目文件未生成: %s", file)
	}
	
	// 检查go.mod内容
	goModContent, err := os.ReadFile(filepath.Join(projectDir, "go.mod"))
	require.NoError(t, err, "无法读取go.mod文件")
	
	expectedModuleDeclaration := "module github.com/username/my-app"
	assert.Contains(t, string(goModContent), expectedModuleDeclaration, "go.mod模块声明不正确")
}

// 辅助函数：创建项目模板文件
func createProjectTemplates(t *testing.T, templatesDir string) error {
	templates := map[string]string{
		"main.go.tmpl": `package main

import (
	"log"
	"{{.Module}}/routers"
)

func main() {
	router := routers.SetupRouter()
	
	log.Println("Starting {{.Name}} server...")
	router.Run(":8080")
}`,
		"go.mod.tmpl": `module {{.Module}}

go 1.19

require (
	github.com/gin-gonic/gin v1.9.0
)
`,
		"README.md.tmpl": `# {{.Name}}

A Gin web application.

## Getting Started

1. Run the application:
   ` + "```" + `
   go run main.go
   ` + "```" + `

2. Visit http://localhost:8080
`,
		"config/config.go.tmpl": `package config

// Config 应用配置
type Config struct {
	AppName string
	Version string
	Port    int
}

// GetConfig 获取默认配置
func GetConfig() *Config {
	return &Config{
		AppName: "{{.Name}}",
		Version: "{{.Version}}",
		Port:    8080,
	}
}
`,
		"routers/router.go.tmpl": `package routers

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to {{.Name}}!",
		})
	})
	
	return router
}
`,
	}
	
	// 创建各个模板文件
	for fileName, content := range templates {
		// 确保父目录存在
		dirPath := filepath.Join(templatesDir, filepath.Dir(fileName))
		if dirPath != templatesDir {
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				return err
			}
		}
		
		filePath := filepath.Join(templatesDir, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return err
		}
	}
	
	// 创建空目录占位符
	dirsToCreate := []string{
		"controllers",
		"models",
		"services",
	}
	
	for _, dir := range dirsToCreate {
		dirPath := filepath.Join(templatesDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
		
		// 创建.gitkeep文件
		gitkeepPath := filepath.Join(dirPath, ".gitkeep.tmpl")
		if err := os.WriteFile(gitkeepPath, []byte(""), 0644); err != nil {
			return err
		}
	}
	
	return nil
}

// 测试初始化项目 - 错误情况
func TestInitProject_Errors(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 创建生成器，但不创建模板目录
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试模板不存在的情况
	err := g.InitProject("test-app", "github.com/username/test-app")
	assert.Error(t, err, "期望在模板不存在时返回错误，但没有")
	
	// 创建模板目录但不包含实际模板
	projectTemplatesDir := filepath.Join(tempDir, "templates", "project")
	err = os.MkdirAll(projectTemplatesDir, 0755)
	require.NoError(t, err, "无法创建项目模板目录")
	
	// 再次测试，应该返回特定的错误
	err = g.InitProject("test-app", "github.com/username/test-app")
	assert.Error(t, err, "期望在缺少关键模板时返回错误，但没有")
} 