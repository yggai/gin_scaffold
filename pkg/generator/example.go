package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/gs/pkg/utils"
)

// ExampleData 示例模板数据
type ExampleData struct {
	Name         string // 示例名称，首字母大写
	PluralName   string // 复数名称，用于列表方法
	ResourceName string // 资源名称，用于URL路径
	Package      string // 项目包名
}

// GenerateExample 生成示例代码
func (g *Generator) GenerateExample(name string, packageName string) error {
	// 格式化名称
	name = formatName(name)
	
	// 准备模板数据
	data := ExampleData{
		Name:         name,
		PluralName:   PluralForm(name),
		ResourceName: strings.ToLower(name) + "s",
		Package:      packageName,
	}
	
	// 确保目录存在
	outputDir := filepath.Join("examples")
	if err := utils.EnsureDir(outputDir); err != nil {
		return fmt.Errorf("无法创建示例目录: %v", err)
	}
	
	// 示例文件路径
	outputFile := filepath.Join(outputDir, strings.ToLower(name)+"_example.go")
	
	// 检查文件是否已存在
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		return fmt.Errorf("示例文件已存在: %s", outputFile)
	}
	
	// 生成示例文件
	templatePath := filepath.Join(g.TemplatesDir, "component", "example", "example.go.tmpl")
	if err := g.GenerateFromTemplate(templatePath, outputFile, data); err != nil {
		return fmt.Errorf("生成示例失败: %v", err)
	}
	
	fmt.Printf("已生成示例文件: %s\n", outputFile)
	return nil
} 