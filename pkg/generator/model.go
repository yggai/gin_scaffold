package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/gs/pkg/utils"
)

// ModelData 模型模板数据
type ModelData struct {
	Name       string // 模型名称，首字母大写
	TableName  string // 表名，全小写
	VarName    string // 变量名称，首字母小写
	Package    string // 项目包名
}

// GenerateModel 生成模型代码
func (g *Generator) GenerateModel(name string, packageName string) error {
	// 格式化名称
	name = formatName(name)
	
	// 准备模板数据
	data := ModelData{
		Name:      name,
		TableName: strings.ToLower(name) + "s",
		VarName:   strings.ToLower(name[:1]) + name[1:],
		Package:   packageName,
	}
	
	// 确保目录存在
	outputDir := filepath.Join("models")
	if err := utils.EnsureDir(outputDir); err != nil {
		return fmt.Errorf("无法创建模型目录: %v", err)
	}
	
	// 模型文件路径
	outputFile := filepath.Join(outputDir, strings.ToLower(name)+".go")
	
	// 检查文件是否已存在
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		return fmt.Errorf("模型文件已存在: %s", outputFile)
	}
	
	// 生成模型文件
	templatePath := filepath.Join(g.TemplatesDir, "component", "model", "model.go.tmpl")
	if err := g.GenerateFromTemplate(templatePath, outputFile, data); err != nil {
		return fmt.Errorf("生成模型失败: %v", err)
	}
	
	fmt.Printf("已生成模型文件: %s\n", outputFile)
	return nil
} 