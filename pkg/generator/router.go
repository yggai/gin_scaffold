package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/gs/pkg/utils"
)

// RouterData 路由模板数据
type RouterData struct {
	Name       string // 控制器名称，首字母大写
	PluralName string // 复数名称，用于列表方法
	Resource   string // 资源名称，用于URL路径
	VarName    string // 变量名称，首字母小写
	Package    string // 项目包名
}

// GenerateRouter 生成路由代码
func (g *Generator) GenerateRouter(name string, packageName string) error {
	// 格式化名称
	name = formatName(name)
	
	// 准备模板数据
	data := RouterData{
		Name:       name,
		PluralName: PluralForm(name),
		Resource:   strings.ToLower(name) + "s",
		VarName:    strings.ToLower(name[:1]) + name[1:],
		Package:    packageName,
	}
	
	// 确保目录存在
	outputDir := filepath.Join("routers")
	if err := utils.EnsureDir(outputDir); err != nil {
		return fmt.Errorf("无法创建路由目录: %v", err)
	}
	
	// 路由文件路径
	outputFile := filepath.Join(outputDir, strings.ToLower(name)+"_router.go")
	
	// 检查文件是否已存在
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		return fmt.Errorf("路由文件已存在: %s", outputFile)
	}
	
	// 生成路由文件
	templatePath := filepath.Join(g.TemplatesDir, "component", "router", "router.go.tmpl")
	if err := g.GenerateFromTemplate(templatePath, outputFile, data); err != nil {
		return fmt.Errorf("生成路由失败: %v", err)
	}
	
	fmt.Printf("已生成路由文件: %s\n", outputFile)
	return nil
} 