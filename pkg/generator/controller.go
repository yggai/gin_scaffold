package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/gs/pkg/utils"
)

// ControllerData 控制器模板数据
type ControllerData struct {
	Name         string // 控制器名称，首字母大写
	PluralName   string // 复数名称，用于列表方法
	ResourceName string // 资源名称，用于URL路径
	VarName      string // 变量名称，首字母小写
	Package      string // 项目包名
}

// GenerateController 生成控制器代码
func (g *Generator) GenerateController(name string, packageName string) error {
	// 格式化名称
	name = formatName(name)
	
	// 准备模板数据
	data := ControllerData{
		Name:         name,
		PluralName:   PluralForm(name),
		ResourceName: strings.ToLower(name) + "s",
		VarName:      strings.ToLower(name[:1]) + name[1:],
		Package:      packageName,
	}
	
	// 确保目录存在
	outputDir := filepath.Join("controllers")
	if err := utils.EnsureDir(outputDir); err != nil {
		return fmt.Errorf("无法创建控制器目录: %v", err)
	}
	
	// 控制器文件路径
	outputFile := filepath.Join(outputDir, strings.ToLower(name)+"_controller.go")
	
	// 检查文件是否已存在
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		return fmt.Errorf("控制器文件已存在: %s", outputFile)
	}
	
	// 生成控制器文件
	templatePath := filepath.Join(g.TemplatesDir, "component", "controller", "controller.go.tmpl")
	if err := g.GenerateFromTemplate(templatePath, outputFile, data); err != nil {
		return fmt.Errorf("生成控制器失败: %v", err)
	}
	
	fmt.Printf("已生成控制器文件: %s\n", outputFile)
	return nil
}

// formatName 格式化名称为Pascal命名（首字母大写）
func formatName(name string) string {
	// 清理名称
	name = utils.SanitizeName(name)
	
	// 首字母大写
	if len(name) > 0 {
		name = strings.ToUpper(name[:1]) + name[1:]
	}
	
	return name
} 