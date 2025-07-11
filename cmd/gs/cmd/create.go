package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yggai/gs/pkg/generator"
)

// createOptions 创建命令选项
type createOptions struct {
	packageName string
	force       bool
}

// NewCreateCmd 创建create命令
func NewCreateCmd() *cobra.Command {
	options := &createOptions{}
	
	cmd := &cobra.Command{
		Use:   "create",
		Short: "创建Gin应用组件",
		Long: `创建Gin应用的各种组件，如控制器、模型、路由等。

例如:
  gs create controller User  # 创建用户控制器
  gs create model User       # 创建用户模型
  gs create router User      # 创建用户路由
  gs create service User     # 创建用户服务
  
您也可以一次性创建多个相关组件:
  gs create resource User    # 创建用户相关的所有组件`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	
	// 共用选项
	options.packageName = getPackageName()
	
	// 添加子命令
	cmd.AddCommand(newCreateControllerCmd(options))
	cmd.AddCommand(newCreateModelCmd(options))
	cmd.AddCommand(newCreateRouterCmd(options))
	cmd.AddCommand(newCreateServiceCmd(options))
	cmd.AddCommand(newCreateResourceCmd(options))
	
	return cmd
}

// newCreateControllerCmd 创建控制器命令
func newCreateControllerCmd(options *createOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "controller [名称]",
		Short: "创建控制器",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			g := generator.NewGenerator(templatesDir)
			if err := g.GenerateController(args[0], options.packageName); err != nil {
				return err
			}
			fmt.Println("控制器创建成功")
			return nil
		},
	}
	
	cmd.Flags().BoolVarP(&options.force, "force", "f", false, "强制创建，覆盖已存在的文件")
	
	return cmd
}

// newCreateModelCmd 创建模型命令
func newCreateModelCmd(options *createOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model [名称]",
		Short: "创建模型",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			g := generator.NewGenerator(templatesDir)
			if err := g.GenerateModel(args[0], options.packageName); err != nil {
				return err
			}
			fmt.Println("模型创建成功")
			return nil
		},
	}
	
	cmd.Flags().BoolVarP(&options.force, "force", "f", false, "强制创建，覆盖已存在的文件")
	
	return cmd
}

// newCreateRouterCmd 创建路由命令
func newCreateRouterCmd(options *createOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "router [名称]",
		Short: "创建路由",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			g := generator.NewGenerator(templatesDir)
			if err := g.GenerateRouter(args[0], options.packageName); err != nil {
				return err
			}
			fmt.Println("路由创建成功")
			return nil
		},
	}
	
	cmd.Flags().BoolVarP(&options.force, "force", "f", false, "强制创建，覆盖已存在的文件")
	
	return cmd
}

// newCreateServiceCmd 创建服务命令
func newCreateServiceCmd(options *createOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service [名称]",
		Short: "创建服务",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			g := generator.NewGenerator(templatesDir)
			if err := g.GenerateService(args[0], options.packageName); err != nil {
				return err
			}
			fmt.Println("服务创建成功")
			return nil
		},
	}
	
	cmd.Flags().BoolVarP(&options.force, "force", "f", false, "强制创建，覆盖已存在的文件")
	
	return cmd
}

// newCreateResourceCmd 创建资源命令（同时创建控制器、模型、路由和服务）
func newCreateResourceCmd(options *createOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resource [名称]",
		Short: "创建完整资源（控制器、模型、路由和服务）",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			g := generator.NewGenerator(templatesDir)
			name := args[0]
			
			// 创建控制器
			if err := g.GenerateController(name, options.packageName); err != nil {
				fmt.Printf("警告: 创建控制器失败: %v\n", err)
			}
			
			// 创建模型
			if err := g.GenerateModel(name, options.packageName); err != nil {
				fmt.Printf("警告: 创建模型失败: %v\n", err)
			}
			
			// 创建路由
			if err := g.GenerateRouter(name, options.packageName); err != nil {
				fmt.Printf("警告: 创建路由失败: %v\n", err)
			}
			
			// 创建服务
			if err := g.GenerateService(name, options.packageName); err != nil {
				fmt.Printf("警告: 创建服务失败: %v\n", err)
			}
			
			fmt.Println("资源创建完成")
			return nil
		},
	}
	
	cmd.Flags().BoolVarP(&options.force, "force", "f", false, "强制创建，覆盖已存在的文件")
	
	return cmd
}

// getPackageName 获取当前Go模块名称
func getPackageName() string {
	// 尝试从go.mod文件获取包名
	if _, err := os.Stat("go.mod"); !os.IsNotExist(err) {
		content, err := os.ReadFile("go.mod")
		if err == nil {
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "module ") {
					return strings.TrimSpace(line[7:])
				}
			}
		}
	}
	
	// 如果无法从go.mod获取，返回默认值
	return "github.com/example/app"
} 