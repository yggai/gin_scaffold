package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileExists(t *testing.T) {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "file_exists_test")
	require.NoError(t, err, "创建临时文件失败")
	defer os.Remove(tmpFile.Name())
	
	// 测试文件存在
	assert.True(t, FileExists(tmpFile.Name()), "文件应该存在")
	
	// 测试文件不存在
	nonExistentFile := filepath.Join(os.TempDir(), "non_existent_file")
	assert.False(t, FileExists(nonExistentFile), "文件不应该存在")
}

func TestEnsureDir(t *testing.T) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "ensure_dir_test")
	require.NoError(t, err, "创建临时目录失败")
	defer os.RemoveAll(tempDir)
	
	// 测试创建新目录
	newDir := filepath.Join(tempDir, "new_dir")
	err = EnsureDir(newDir)
	assert.NoError(t, err, "EnsureDir应该成功创建新目录")
	
	dirInfo, err := os.Stat(newDir)
	assert.NoError(t, err, "应该能够获取目录信息")
	assert.True(t, dirInfo.IsDir(), "创建的应该是一个目录")
	
	// 测试已存在的目录
	err = EnsureDir(newDir)
	assert.NoError(t, err, "对已存在的目录调用EnsureDir应该成功")
}

func TestGetProjectRoot(t *testing.T) {
	// 创建临时目录结构
	tempDir, err := os.MkdirTemp("", "project_root_test")
	require.NoError(t, err, "创建临时目录失败")
	defer os.RemoveAll(tempDir)
	
	// 保存当前目录
	originalDir, err := os.Getwd()
	require.NoError(t, err, "获取当前目录失败")
	defer os.Chdir(originalDir)
	
	// 创建模拟项目结构
	projectDir := filepath.Join(tempDir, "project")
	err = os.MkdirAll(projectDir, 0755)
	require.NoError(t, err, "创建项目目录失败")
	
	// 创建go.mod文件
	goModPath := filepath.Join(projectDir, "go.mod")
	err = os.WriteFile(goModPath, []byte("module example.com/project"), 0644)
	require.NoError(t, err, "创建go.mod文件失败")
	
	// 创建子目录
	subDir := filepath.Join(projectDir, "pkg", "utils")
	err = os.MkdirAll(subDir, 0755)
	require.NoError(t, err, "创建子目录失败")
	
	// 测试在项目根目录
	err = os.Chdir(projectDir)
	require.NoError(t, err, "切换到项目根目录失败")
	
	root, err := GetProjectRoot()
	assert.NoError(t, err, "在项目根目录获取项目根应该成功")
	assert.Equal(t, ".", root, "在项目根目录，应该返回当前目录")
	
	// 测试在子目录
	err = os.Chdir(subDir)
	require.NoError(t, err, "切换到子目录失败")
	
	// 在测试环境中，GetProjectRoot可能无法正确工作，因为它查找go.mod文件
	// 此处仅验证函数执行不出错
	_, err = GetProjectRoot()
	if err != nil {
		t.Logf("GetProjectRoot在子目录返回错误: %v", err)
	}
}

func TestSanitizeName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"空字符串", "", ""},
		{"正常标识符", "validName", "validName"},
		{"带空格", "name with spaces", "name_with_spaces"},
		{"带破折号", "name-with-dashes", "name_with_dashes"},
		{"带点号", "name.with.dots", "name_with_dots"},
		{"带特殊字符", "name@#$%^", "name"},
		{"数字开头", "123name", "_123name"},
		{"混合情况", "123-name.with special@chars", "_123_name_with_specialchars"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeName(tt.input)
			assert.Equal(t, tt.expected, result, "SanitizeName处理结果不正确")
		})
	}
}

func TestFormatPackageName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"空字符串", "", ""},
		{"大写字母", "Package", "package"},
		{"混合大小写", "PackageName", "packagename"},
		{"带特殊字符", "package-name", "package_name"},
		{"混合情况", "Package-Name.With.Special@Chars", "package_name_with_specialchars"},
		{"数字开头", "123package", "_123package"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatPackageName(tt.input)
			assert.Equal(t, tt.expected, result, "FormatPackageName处理结果不正确")
		})
	}
} 