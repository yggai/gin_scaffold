package generator

import (
	"fmt"
)

// GenerateFeature 生成完整功能代码，包含模型、服务、控制器、路由等
func (g *Generator) GenerateFeature(name string, packageName string) error {
	// 格式化名称
	name = formatName(name)
	
	// 生成模型
	if err := g.GenerateModel(name, packageName); err != nil {
		return fmt.Errorf("生成模型失败: %v", err)
	}
	
	// 生成服务
	if err := g.GenerateService(name, packageName); err != nil {
		return fmt.Errorf("生成服务失败: %v", err)
	}
	
	// 生成控制器
	if err := g.GenerateController(name, packageName); err != nil {
		return fmt.Errorf("生成控制器失败: %v", err)
	}
	
	// 生成路由
	if err := g.GenerateRoute(name, packageName); err != nil {
		return fmt.Errorf("生成路由失败: %v", err)
	}
	
	// 生成测试
	if err := g.GenerateTest(name, packageName); err != nil {
		return fmt.Errorf("生成测试失败: %v", err)
	}
	
	// 生成示例
	if err := g.GenerateExample(name, packageName); err != nil {
		return fmt.Errorf("生成示例失败: %v", err)
	}
	
	fmt.Printf("已成功生成 %s 的完整功能代码\n", name)
	return nil
} 