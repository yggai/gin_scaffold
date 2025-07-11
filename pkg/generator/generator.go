package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Generator 代码生成器
type Generator struct {
	TemplatesDir string
}

// NewGenerator 创建一个新的代码生成器
func NewGenerator(templatesDir string) *Generator {
	return &Generator{
		TemplatesDir: templatesDir,
	}
}

// GenerateFromTemplate 从模板生成代码
func (g *Generator) GenerateFromTemplate(templateName string, outputPath string, data interface{}) error {
	// 检查模板文件是否存在
	templatePath := templateName
	if !filepath.IsAbs(templatePath) {
		templatePath = filepath.Join(g.TemplatesDir, templateName)
	}
	
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return fmt.Errorf("模板文件不存在: %s", templatePath)
	}

	// 读取模板内容
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("无法读取模板文件: %v", err)
	}

	// 创建输出目录
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("无法创建输出目录: %v", err)
	}

	// 解析模板
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("无法解析模板: %v", err)
	}

	// 生成文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outputFile.Close()

	// 执行模板
	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("模板执行失败: %v", err)
	}

	return nil
}

// CapitalizeFirst 将字符串的第一个字母大写
func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// LowercaseFirst 将字符串的第一个字母小写
func LowercaseFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// PluralForm 返回名词的复数形式（简单处理）
func PluralForm(s string) string {
	if s == "" {
		return ""
	}
	
	// 简单的复数形式处理
	if strings.HasSuffix(s, "y") {
		return s[:len(s)-1] + "ies"
	}
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") || 
	   strings.HasSuffix(s, "z") || strings.HasSuffix(s, "ch") || strings.HasSuffix(s, "sh") {
		return s + "es"
	}
	
	return s + "s"
} 