package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试辅助函数：创建临时目录
func createTempDir(t *testing.T) string {
	dir, err := os.MkdirTemp("", "generator-test-")
	require.NoError(t, err, "无法创建临时测试目录")
	return dir
}

// 测试辅助函数：创建临时文件
func createTempFile(t *testing.T, dir, name, content string) string {
	path := filepath.Join(dir, name)
	err := os.WriteFile(path, []byte(content), 0644)
	require.NoError(t, err, "无法创建临时测试文件")
	return path
}

// 测试辅助函数：清理临时目录
func cleanupTempDir(t *testing.T, dir string) {
	if dir != "" {
		err := os.RemoveAll(dir)
		if err != nil {
			t.Logf("警告: 清理临时目录失败: %v", err)
		}
	}
}

// 测试NewGenerator函数
func TestNewGenerator(t *testing.T) {
	templatesDir := "/path/to/templates"
	g := NewGenerator(templatesDir)
	
	assert.NotNil(t, g, "NewGenerator返回了nil")
	assert.Equal(t, templatesDir, g.TemplatesDir, "模板目录设置不正确")
}

// 测试GenerateFromTemplate函数
func TestGenerateFromTemplate(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 创建测试模板
	templateContent := `package example

type {{.Name}} struct {
	ID string
	Value string
}
`
	templatePath := createTempFile(t, tempDir, "test.tmpl", templateContent)
	
	// 创建生成器
	g := NewGenerator(tempDir)
	
	// 准备输出路径和测试数据
	outputPath := filepath.Join(tempDir, "output.go")
	data := struct {
		Name string
	}{
		Name: "TestStruct",
	}
	
	// 测试生成
	err := g.GenerateFromTemplate(filepath.Base(templatePath), outputPath, data)
	require.NoError(t, err, "生成文件失败")
	
	// 验证生成的文件内容
	content, err := os.ReadFile(outputPath)
	require.NoError(t, err, "无法读取生成的文件")
	
	expectedContent := `package example

type TestStruct struct {
	ID string
	Value string
}
`
	assert.Equal(t, expectedContent, string(content), "生成的内容不匹配")
}

// 测试GenerateFromTemplate函数 - 错误情况
func TestGenerateFromTemplate_Errors(t *testing.T) {
	// 创建测试环境
	tempDir := createTempDir(t)
	defer cleanupTempDir(t, tempDir)
	
	// 创建生成器
	g := NewGenerator(tempDir)
	
	// 测试不存在的模板
	err := g.GenerateFromTemplate("non_existent.tmpl", "output.go", nil)
	assert.Error(t, err, "期望对不存在的模板返回错误，但没有")
	
	// 测试无效的模板
	invalidTemplatePath := createTempFile(t, tempDir, "invalid.tmpl", "{{ .Name }")
	err = g.GenerateFromTemplate(filepath.Base(invalidTemplatePath), filepath.Join(tempDir, "output.go"), nil)
	assert.Error(t, err, "期望对无效的模板返回错误，但没有")
	
	// 测试无效的输出路径
	validTemplatePath := createTempFile(t, tempDir, "valid.tmpl", "content")
	err = g.GenerateFromTemplate(filepath.Base(validTemplatePath), "/invalid/path/output.go", nil)
	assert.Error(t, err, "期望对无效的输出路径返回错误，但没有")
}

// 测试formatName函数
func TestFormatName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"空字符串", "", ""},
		{"小写字符", "test", "Test"},
		{"已大写字符", "Test", "Test"},
		{"带下划线", "test_name", "Test_name"},
		{"带空格", "test name", "Test_name"},
		{"带破折号", "test-name", "Test_name"},
		{"数字开头", "123test", "_123test"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatName(tt.input)
			assert.Equal(t, tt.expected, result, "formatName处理结果不正确")
		})
	}
}

// 测试PluralForm函数
func TestPluralForm(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"空字符串", "", ""},
		{"常规名词", "book", "books"},
		{"以s结尾", "class", "classes"},
		{"以x结尾", "box", "boxes"},
		{"以z结尾", "quiz", "quizzes"},
		{"以ch结尾", "watch", "watches"},
		{"以sh结尾", "dish", "dishes"},
		{"以y结尾", "city", "cities"},
		{"已是复数形式", "books", "bookses"}, // 简单实现，不处理已是复数的情况
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PluralForm(tt.input)
			assert.Equal(t, tt.expected, result, "复数形式转换不正确")
		})
	}
}

// 测试CapitalizeFirst和LowercaseFirst函数
func TestStringHelpers(t *testing.T) {
	tests := []struct {
		input            string
		expectedCapital  string
		expectedLower    string
	}{
		{"", "", ""},
		{"test", "Test", "test"},
		{"Test", "Test", "test"},
		{"TEST", "TEST", "tEST"},
		{"t", "T", "t"},
		{"T", "T", "t"},
	}
	
	for _, tt := range tests {
		t.Run("CapitalizeFirst_"+tt.input, func(t *testing.T) {
			result := CapitalizeFirst(tt.input)
			assert.Equal(t, tt.expectedCapital, result, "首字母大写转换不正确")
		})
		
		t.Run("LowercaseFirst_"+tt.input, func(t *testing.T) {
			result := LowercaseFirst(tt.input)
			assert.Equal(t, tt.expectedLower, result, "首字母小写转换不正确")
		})
	}
} 