package gs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yggai/gs/pkg/generator"
)

// 获取默认项目包名
func getDefaultPackage() string {
	// 尝试从go.mod获取包名
	modContent, err := os.ReadFile("go.mod")
	if err == nil {
		lines := strings.Split(string(modContent), "\n")
		if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
			return strings.TrimSpace(strings.TrimPrefix(lines[0], "module "))
		}
	}
	
	// 如果无法从go.mod获取，则使用当前目录名
	dir, err := os.Getwd()
	if err == nil {
		return filepath.Base(dir)
	}
	
	// 默认包名
	return "myapp"
}

// 获取模板目录
func getTemplatesDir() (string, error) {
	// 先尝试直接在当前目录查找
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("无法获取当前工作目录: %v", err)
	}
	
	// 检查当前目录下的templates目录
	currentTemplatesDir := filepath.Join(currentDir, "templates")
	if _, err := os.Stat(currentTemplatesDir); !os.IsNotExist(err) {
		return currentTemplatesDir, nil
	}
	
	// 查找项目根目录的templates目录
	projectRoot := currentDir
	for i := 0; i < 3; i++ { // 向上最多查找3层
		templatesDir := filepath.Join(projectRoot, "templates")
		if _, err := os.Stat(templatesDir); !os.IsNotExist(err) {
			return templatesDir, nil
		}
		parentDir := filepath.Dir(projectRoot)
		if parentDir == projectRoot {
			break
		}
		projectRoot = parentDir
	}
	
	// 如果在项目目录找不到，则尝试可执行文件所在目录
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("无法获取可执行文件路径: %v", err)
	}
	execDir := filepath.Dir(execPath)
	execTemplatesDir := filepath.Join(execDir, "templates")
	if _, err := os.Stat(execTemplatesDir); !os.IsNotExist(err) {
		return execTemplatesDir, nil
	}
	
	// 所有位置都找不到，返回错误
	return "", fmt.Errorf("无法找到模板目录。请确保templates目录存在于当前目录或项目根目录")
}

// createCmd 生成各种组件
var createCmd = &cobra.Command{
	Use:   "create [组件类型] [名称]",
	Short: "创建Gin项目组件",
	Long: `创建Gin项目组件，包括控制器、路由、模型、服务等。
可用的组件类型:
  controller  - 创建控制器
  route       - 创建路由
  model       - 创建数据模型
  service     - 创建服务
  example     - 创建示例代码
  test        - 创建测试代码
  feature     - 创建完整功能集`,
	Run: func(cmd *cobra.Command, args []string) {
		// 如果没有提供足够的参数，显示帮助信息
		if len(args) < 2 {
			cmd.Help()
			return
		}
		
		componentType := strings.ToLower(args[0])
		componentName := args[1]
		
		// 获取模板目录
		templatesDir, err := getTemplatesDir()
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}
		
		// 创建生成器
		g := generator.NewGenerator(templatesDir)
		
		// 获取项目包名
		packageName, _ := cmd.Flags().GetString("package")
		if packageName == "" {
			packageName = getDefaultPackage()
		}
		
		switch componentType {
		case "controller":
			if err := g.GenerateController(componentName, packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		case "route":
			if err := g.GenerateRoute(componentName, packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		case "model":
			if err := g.GenerateModel(componentName, packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		case "service":
			if err := g.GenerateService(componentName, packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		case "example":
			if err := g.GenerateExample(componentName, packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		case "test":
			if err := g.GenerateTest(componentName, packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		case "feature":
			if err := g.GenerateFeature(componentName, packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		default:
			fmt.Printf("错误：不支持的组件类型 '%s'\n", componentType)
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	
	// 为create命令添加选项
	createCmd.PersistentFlags().String("package", "", "项目包名(默认从go.mod获取)")
	
	// 为create命令添加子命令
	createCmd.AddCommand(createControllerCmd())
	createCmd.AddCommand(createRouteCmd())
	createCmd.AddCommand(createModelCmd())
	createCmd.AddCommand(createServiceCmd())
	createCmd.AddCommand(createExampleCmd())
	createCmd.AddCommand(createTestCmd())
	createCmd.AddCommand(createFeatureCmd())
}

// 创建控制器命令
func createControllerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "controller [名称]",
		Short: "创建控制器",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 获取模板目录
			templatesDir, err := getTemplatesDir()
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 获取项目包名
			packageName, _ := cmd.Flags().GetString("package")
			if packageName == "" {
				packageName = getDefaultPackage()
			}
			
			// 生成控制器
			if err := g.GenerateController(args[0], packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		},
	}
}

// 创建路由命令
func createRouteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "route [名称]",
		Short: "创建路由",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 获取模板目录
			templatesDir, err := getTemplatesDir()
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 获取项目包名
			packageName, _ := cmd.Flags().GetString("package")
			if packageName == "" {
				packageName = getDefaultPackage()
			}
			
			// 生成路由
			if err := g.GenerateRoute(args[0], packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		},
	}
}

// 创建模型命令
func createModelCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "model [名称]",
		Short: "创建数据模型",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 获取模板目录
			templatesDir, err := getTemplatesDir()
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 获取项目包名
			packageName, _ := cmd.Flags().GetString("package")
			if packageName == "" {
				packageName = getDefaultPackage()
			}
			
			// 生成模型
			if err := g.GenerateModel(args[0], packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		},
	}
}

// 创建服务命令
func createServiceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "service [名称]",
		Short: "创建服务",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 获取模板目录
			templatesDir, err := getTemplatesDir()
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 获取项目包名
			packageName, _ := cmd.Flags().GetString("package")
			if packageName == "" {
				packageName = getDefaultPackage()
			}
			
			// 生成服务
			if err := g.GenerateService(args[0], packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		},
	}
}

// 创建示例命令
func createExampleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "example [名称]",
		Short: "创建示例代码",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 获取模板目录
			templatesDir, err := getTemplatesDir()
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 获取项目包名
			packageName, _ := cmd.Flags().GetString("package")
			if packageName == "" {
				packageName = getDefaultPackage()
			}
			
			// 生成示例
			if err := g.GenerateExample(args[0], packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		},
	}
}

// 创建测试命令
func createTestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test [名称]",
		Short: "创建测试代码",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 获取模板目录
			templatesDir, err := getTemplatesDir()
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 获取项目包名
			packageName, _ := cmd.Flags().GetString("package")
			if packageName == "" {
				packageName = getDefaultPackage()
			}
			
			// 生成测试
			if err := g.GenerateTest(args[0], packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		},
	}
}

// 创建完整功能命令
func createFeatureCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "feature [名称]",
		Short: "创建完整功能",
		Long:  "创建完整功能集，包括模型、服务、控制器、路由、示例代码和测试代码",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 获取模板目录
			templatesDir, err := getTemplatesDir()
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			
			// 创建生成器
			g := generator.NewGenerator(templatesDir)
			
			// 获取项目包名
			packageName, _ := cmd.Flags().GetString("package")
			if packageName == "" {
				packageName = getDefaultPackage()
			}
			
			// 生成完整功能
			if err := g.GenerateFeature(args[0], packageName); err != nil {
				fmt.Printf("错误: %v\n", err)
			}
		},
	}
} 