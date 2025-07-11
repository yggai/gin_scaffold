package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yggai/gs/pkg/generator"
)

// initOptions 初始化命令选项
type initOptions struct {
	moduleName string
	force      bool
}

// NewInitCmd 创建初始化命令
func NewInitCmd() *cobra.Command {
	options := &initOptions{}
	
	cmd := &cobra.Command{
		Use:   "init [项目名称]",
		Short: "初始化一个新的Gin应用",
		Long: `初始化一个新的Gin Web应用程序，包括基本的项目结构和配置文件。

例如:
  gs init myapp                       # 在当前目录下创建新项目
  gs init myapp --module github.com/username/myapp  # 指定Go模块名称`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 初始化项目
			if err := g.InitProject(projectName, options.moduleName); err != nil {
				return fmt.Errorf("项目初始化失败: %v", err)
			}
			
			return nil
		},
	}
	
	// 添加命令选项
	cmd.Flags().StringVarP(&options.moduleName, "module", "m", "", "Go模块名称 (默认与项目名称相同)")
	cmd.Flags().BoolVarP(&options.force, "force", "f", false, "强制初始化，即使目标目录已存在")
	
	return cmd
} 