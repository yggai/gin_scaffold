package generator

import (
	"os"
	"path/filepath"
	"testing"
)

// 测试路由数据结构
func TestRouterData(t *testing.T) {
	data := RouterData{
		Name:       "User",
		PluralName: "Users",
		Resource:   "users",
		VarName:    "user",
		Package:    "myapp",
	}
	
	if data.Name != "User" {
		t.Errorf("Name字段设置不正确，期望: User, 实际: %s", data.Name)
	}
	
	if data.PluralName != "Users" {
		t.Errorf("PluralName字段设置不正确，期望: Users, 实际: %s", data.PluralName)
	}
	
	if data.Resource != "users" {
		t.Errorf("Resource字段设置不正确，期望: users, 实际: %s", data.Resource)
	}
	
	if data.VarName != "user" {
		t.Errorf("VarName字段设置不正确，期望: user, 实际: %s", data.VarName)
	}
	
	if data.Package != "myapp" {
		t.Errorf("Package字段设置不正确，期望: myapp, 实际: %s", data.Package)
	}
}

// 测试生成路由
func TestGenerateRouter(t *testing.T) {
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
	templatesDir := filepath.Join(tempDir, "templates", "component", "router")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("无法创建模板目录: %v", err)
	}
	
	// 创建测试模板
	templateContent := `package routers

import (
	"github.com/gin-gonic/gin"
	"{{.Package}}/controllers"
)

// Setup{{.Name}}Routes 设置{{.Name}}相关的路由
func Setup{{.Name}}Routes(router *gin.Engine) {
	{{.VarName}}Controller := &controllers.{{.Name}}Controller{}
	
	// 创建{{.Name}}资源路由组
	{{.Resource}}Group := router.Group("/api/{{.Resource}}")
	{
		{{.Resource}}Group.GET("", {{.VarName}}Controller.Get{{.PluralName}})
		{{.Resource}}Group.POST("", {{.VarName}}Controller.Create{{.Name}})
		{{.Resource}}Group.GET("/:id", {{.VarName}}Controller.Get{{.Name}})
		{{.Resource}}Group.PUT("/:id", {{.VarName}}Controller.Update{{.Name}})
		{{.Resource}}Group.DELETE("/:id", {{.VarName}}Controller.Delete{{.Name}})
	}
}
`
	templatePath := filepath.Join(templatesDir, "router.go.tmpl")
	if err := os.WriteFile(templatePath, []byte(templateContent), 0644); err != nil {
		t.Fatalf("无法创建测试模板文件: %v", err)
	}
	
	// 创建生成器
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试生成路由
	err = g.GenerateRouter("User", "myapp")
	if err != nil {
		t.Fatalf("生成路由失败: %v", err)
	}
	
	// 验证路由文件是否已生成
	routerPath := filepath.Join(tempDir, "routers", "user_router.go")
	if _, err := os.Stat(routerPath); os.IsNotExist(err) {
		t.Fatalf("路由文件未生成: %s", routerPath)
	}
	
	// 验证路由内容
	content, err := os.ReadFile(routerPath)
	if err != nil {
		t.Fatalf("无法读取生成的路由文件: %v", err)
	}
	
	// 检查内容是否符合预期
	expectedContent := `package routers

import (
	"github.com/gin-gonic/gin"
	"myapp/controllers"
)

// SetupUserRoutes 设置User相关的路由
func SetupUserRoutes(router *gin.Engine) {
	userController := &controllers.UserController{}
	
	// 创建User资源路由组
	usersGroup := router.Group("/api/users")
	{
		usersGroup.GET("", userController.GetUsers)
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("/:id", userController.GetUser)
		usersGroup.PUT("/:id", userController.UpdateUser)
		usersGroup.DELETE("/:id", userController.DeleteUser)
	}
}
`
	if string(content) != expectedContent {
		t.Errorf("生成的路由内容不符合预期\n期望:\n%s\n实际:\n%s", expectedContent, string(content))
	}
}

// 测试生成路由 - 错误情况
func TestGenerateRouter_Errors(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 创建生成器，但不创建模板目录
	g := NewGenerator(filepath.Join(tempDir, "templates"))
	
	// 测试模板不存在的情况
	err := g.GenerateRouter("User", "myapp")
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
	templatesDir := filepath.Join(tempDir, "templates", "component", "router")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("无法创建模板目录: %v", err)
	}
	
	// 创建测试模板
	templatePath := filepath.Join(templatesDir, "router.go.tmpl")
	if err := os.WriteFile(templatePath, []byte("template content"), 0644); err != nil {
		t.Fatalf("无法创建测试模板文件: %v", err)
	}
	
	// 创建routers目录和已存在的路由文件
	routersDir := filepath.Join(tempDir, "routers")
	if err := os.MkdirAll(routersDir, 0755); err != nil {
		t.Fatalf("无法创建routers目录: %v", err)
	}
	
	// 创建已存在的路由文件
	routerPath := filepath.Join(routersDir, "user_router.go")
	if err := os.WriteFile(routerPath, []byte("already exists"), 0644); err != nil {
		t.Fatalf("无法创建已存在的路由文件: %v", err)
	}
	
	// 测试路由文件已存在的情况
	err = g.GenerateRouter("User", "myapp")
	if err == nil {
		t.Error("期望在路由文件已存在时返回错误，但没有")
	}
} 