package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version      = "v0.1.0" // 版本号
	templatesDir string     // 模板目录
)

// NewRootCmd 创建根命令
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gs",
		Short: "Gin脚手架工具",
		Long: `Gin脚手架工具 (gs) 是一个用于快速生成Gin Web应用程序的命令行工具。
它可以帮助你创建控制器、模型、路由等组件。`,
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			if showVersion, _ := cmd.Flags().GetBool("version"); showVersion {
				fmt.Printf("gs版本 %s\n", version)
				return nil
			}
			return cmd.Help()
		},
	}

	rootCmd.Flags().BoolP("version", "v", false, "显示版本信息")

	// 初始化模板目录
	initTemplatesDir()

	// 添加子命令
	rootCmd.AddCommand(NewInitCmd())
	rootCmd.AddCommand(NewCreateCmd())

	return rootCmd
}

// Execute 执行根命令
func Execute() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initTemplatesDir 初始化模板目录路径
func initTemplatesDir() {
	// 首先尝试从环境变量获取模板目录
	templatesDir = os.Getenv("GS_TEMPLATES_DIR")
	if templatesDir != "" {
		return
	}

	// 尝试从可执行文件所在目录获取
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		possibleTemplateDirs := []string{
			filepath.Join(exeDir, "templates"),
			filepath.Join(exeDir, "..", "templates"),
			filepath.Join(exeDir, "..", "..", "templates"),
		}

		for _, dir := range possibleTemplateDirs {
			if _, err := os.Stat(dir); !os.IsNotExist(err) {
				templatesDir = dir
				return
			}
		}
	}

	// 尝试从当前目录获取
	cwd, err := os.Getwd()
	if err == nil {
		possibleTemplateDirs := []string{
			filepath.Join(cwd, "templates"),
			filepath.Join(cwd, "..", "templates"),
		}

		for _, dir := range possibleTemplateDirs {
			if _, err := os.Stat(dir); !os.IsNotExist(err) {
				templatesDir = dir
				return
			}
		}
	}

	// 如果是Windows系统，可能需要处理特殊路径
	if runtime.GOOS == "windows" {
		// 尝试获取GOPATH
		gopath := os.Getenv("GOPATH")
		if gopath != "" {
			possibleDir := filepath.Join(gopath, "src", "github.com", "yggai", "gs", "templates")
			if _, err := os.Stat(possibleDir); !os.IsNotExist(err) {
				templatesDir = possibleDir
				return
			}
		}
	}

	// 最后，使用相对于当前目录的默认路径
	templatesDir = filepath.Join("templates")
} 