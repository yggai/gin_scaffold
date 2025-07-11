package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/gs/pkg/utils"
)

// TestData 测试模板数据
type TestData struct {
	Name         string // 测试名称，首字母大写
	PluralName   string // 复数名称，用于列表方法
	ResourceName string // 资源名称，用于URL路径
	Package      string // 项目包名
}

// GenerateTest 生成测试代码
func (g *Generator) GenerateTest(name string, packageName string) error {
	// 格式化名称
	name = formatName(name)
	
	// 准备模板数据
	data := TestData{
		Name:         name,
		PluralName:   PluralForm(name),
		ResourceName: strings.ToLower(name) + "s",
		Package:      packageName,
	}
	
	// 确保目录存在
	outputDir := filepath.Join("tests")
	if err := utils.EnsureDir(outputDir); err != nil {
		return fmt.Errorf("无法创建测试目录: %v", err)
	}
	
	// 测试文件路径
	outputFile := filepath.Join(outputDir, strings.ToLower(name)+"_test.go")
	
	// 检查文件是否已存在
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		return fmt.Errorf("测试文件已存在: %s", outputFile)
	}
	
	// 生成测试文件
	templatePath := filepath.Join(g.TemplatesDir, "component", "test", "test.go.tmpl")
	if err := g.GenerateFromTemplate(templatePath, outputFile, data); err != nil {
		return fmt.Errorf("生成测试失败: %v", err)
	}
	
	fmt.Printf("已生成测试文件: %s\n", outputFile)
	return nil
} 