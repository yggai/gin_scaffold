package gs

import (
	"github.com/spf13/cobra"
)

var (
	// 版本信息，将在编译时通过 -ldflags 设置
	version = "v0.1.0"
	
	// 根命令
	rootCmd = &cobra.Command{
		Use:   "gs",
		Short: "Gin脚手架命令行工具",
		Long: `gs是一个轻量级命令行工具，用于快速生成基于Gin框架的Go项目代码。
可以生成项目结构、控制器、路由、模型、服务等各种组件代码。`,
		Version: version,
	}
)

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 初始化命令行选项
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "启用详细输出模式")
	
	// 添加子命令
	// 这些命令将在各自的文件中实现
} 