package gs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd 初始化新的Gin项目
var initCmd = &cobra.Command{
	Use:   "init [项目名称]",
	Short: "创建一个新的Gin项目",
	Long:  `创建一个新的Gin项目，包含标准的目录结构和基础配置文件。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		// 验证项目名称
		if !isValidProjectName(projectName) {
			fmt.Println("错误：项目名称只能包含字母、数字、下划线和连字符")
			return
		}
		
		// 检查目录是否已存在
		if _, err := os.Stat(projectName); !os.IsNotExist(err) {
			fmt.Printf("错误：目录 '%s' 已存在\n", projectName)
			return
		}
		
		// 创建项目目录
		fmt.Printf("创建项目 '%s'\n", projectName)
		if err := os.MkdirAll(projectName, 0755); err != nil {
			fmt.Printf("错误：无法创建项目目录: %v\n", err)
			return
		}
		
		// 创建项目结构
		directories := []string{
			"config",
			"controllers",
			"middlewares",
			"models",
			"services",
			"routes",
			"utils",
			"tests",
			"examples",
		}
		
		for _, dir := range directories {
			path := filepath.Join(projectName, dir)
			if err := os.MkdirAll(path, 0755); err != nil {
				fmt.Printf("错误：创建目录失败 %s: %v\n", path, err)
				return
			}
		}
		
		// 创建go.mod文件
		goModPath := filepath.Join(projectName, "go.mod")
		modContent := fmt.Sprintf("module %s\n\ngo 1.16\n", projectName)
		if err := os.WriteFile(goModPath, []byte(modContent), 0644); err != nil {
			fmt.Printf("错误：无法创建go.mod文件: %v\n", err)
			return
		}
		
		// 创建main.go文件
		mainPath := filepath.Join(projectName, "main.go")
		mainContent := `package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 注册路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 启动服务器
	fmt.Println("服务器启动在 :8080 端口...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
`
		if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
			fmt.Printf("错误：无法创建main.go文件: %v\n", err)
			return
		}
		
		fmt.Println("项目创建成功！")
		fmt.Println("安装依赖:")
		fmt.Printf("cd %s && go mod tidy\n", projectName)
	},
}

// isValidProjectName 检查项目名称是否有效
func isValidProjectName(name string) bool {
	// 项目名称只允许字母、数字、下划线和连字符
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '_' || char == '-') {
			return false
		}
	}
	return true
}

func init() {
	rootCmd.AddCommand(initCmd)
} 