package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ProjectData 项目模板数据
type ProjectData struct {
	Name    string // 项目名称
	Module  string // Go模块名称
	Version string // 版本号
}

// InitProject 初始化项目
func (g *Generator) InitProject(name string, moduleName string) error {
	// 验证项目名称
	if name == "" {
		return fmt.Errorf("项目名称不能为空")
	}
	
	// 如果未提供模块名称，使用项目名称
	if moduleName == "" {
		moduleName = name
	}
	
	// 准备模板数据
	data := ProjectData{
		Name:    name,
		Module:  moduleName,
		Version: "v0.1.0",
	}
	
	// 项目目录路径
	projectDir := name
	if _, err := os.Stat(projectDir); !os.IsNotExist(err) && projectDir != "." {
		return fmt.Errorf("目录已存在: %s", projectDir)
	}
	
	// 确保项目目录存在
	if projectDir != "." {
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			return fmt.Errorf("无法创建项目目录: %v", err)
		}
	}
	
	// 获取项目模板目录路径
	templatesDir := filepath.Join(g.TemplatesDir, "project")
	
	// 递归遍历模板目录并生成项目文件
	return g.generateProjectFiles(templatesDir, projectDir, data)
}

// generateProjectFiles 递归生成项目文件
func (g *Generator) generateProjectFiles(templatesDir, outputDir string, data ProjectData) error {
	// 获取模板目录中的所有文件和子目录
	entries, err := os.ReadDir(templatesDir)
	if err != nil {
		return fmt.Errorf("无法读取模板目录: %v", err)
	}
	
	// 遍历每个条目
	for _, entry := range entries {
		templatePath := filepath.Join(templatesDir, entry.Name())
		
		// 计算输出路径，移除.tmpl后缀
		outputName := strings.TrimSuffix(entry.Name(), ".tmpl")
		outputPath := filepath.Join(outputDir, outputName)
		
		if entry.IsDir() {
			// 如果是目录，则创建对应的输出目录并递归处理
			if err := os.MkdirAll(outputPath, 0755); err != nil {
				return fmt.Errorf("无法创建目录: %v", err)
			}
			
			if err := g.generateProjectFiles(templatePath, outputPath, data); err != nil {
				return err
			}
		} else {
			// 如果是文件，则生成项目文件
			if filepath.Ext(templatePath) == ".tmpl" {
				// 如果是模板文件，使用模板引擎渲染
				if err := g.GenerateFromTemplate(templatePath, outputPath, data); err != nil {
					return fmt.Errorf("生成文件失败 %s: %v", outputPath, err)
				}
				fmt.Printf("已生成文件: %s\n", outputPath)
			} else {
				// 否则直接复制文件
				content, err := os.ReadFile(templatePath)
				if err != nil {
					return fmt.Errorf("无法读取文件: %v", err)
				}
				
				if err := os.WriteFile(outputPath, content, 0644); err != nil {
					return fmt.Errorf("无法写入文件: %v", err)
				}
				fmt.Printf("已复制文件: %s\n", outputPath)
			}
		}
	}
	
	fmt.Printf("项目 %s 初始化成功！\n", data.Name)
	return nil
} 