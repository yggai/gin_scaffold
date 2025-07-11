package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/gs/pkg/utils"
)

// ServiceData 服务模板数据
type ServiceData struct {
	Name      string // 服务名称，首字母大写
	VarName   string // 变量名称，首字母小写
	Package   string // 项目包名
}

// GenerateService 生成服务代码
func (g *Generator) GenerateService(name string, packageName string) error {
	// 格式化名称
	name = formatName(name)
	
	// 准备模板数据
	data := ServiceData{
		Name:     name,
		VarName:  strings.ToLower(name[:1]) + name[1:],
		Package:  packageName,
	}
	
	// 确保目录存在
	outputDir := filepath.Join("services")
	if err := utils.EnsureDir(outputDir); err != nil {
		return fmt.Errorf("无法创建服务目录: %v", err)
	}
	
	// 服务文件路径
	outputFile := filepath.Join(outputDir, strings.ToLower(name)+"_service.go")
	
	// 检查文件是否已存在
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		return fmt.Errorf("服务文件已存在: %s", outputFile)
	}
	
	// 生成服务文件
	templatePath := filepath.Join(g.TemplatesDir, "component", "service", "service.go.tmpl")
	if err := g.GenerateFromTemplate(templatePath, outputFile, data); err != nil {
		return fmt.Errorf("生成服务失败: %v", err)
	}
	
	fmt.Printf("已生成服务文件: %s\n", outputFile)
	return nil
} 