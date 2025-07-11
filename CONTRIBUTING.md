# 贡献指南

感谢您对Gin脚手架工具(GS)的关注！我们欢迎并鼓励社区参与贡献。本文档将指导您如何参与本项目的开发和改进。

## 行为准则

请保持专业、尊重和包容的态度。我们希望维护一个友好、开放的社区环境。

## 贡献方式

您可以通过多种方式为项目做出贡献：

1. **代码贡献**：修复bug或添加新功能
2. **文档改进**：完善现有文档或添加缺失内容
3. **问题报告**：提交bug或建议改进的功能
4. **测试**：增加测试覆盖率或改进现有测试
5. **代码审查**：审查其他贡献者的Pull Request

## 贡献流程

### 1. 准备工作

1. Fork本项目到您的GitHub账户
2. Clone您fork的仓库到本地
   ```bash
   git clone https://github.com/YOUR-USERNAME/gs.git
   cd gs
   ```
3. 添加原始仓库作为上游远程仓库
   ```bash
   git remote add upstream https://github.com/original-owner/gs.git
   ```

### 2. 创建分支

为您的贡献创建一个新分支：

```bash
git checkout -b feature/my-feature  # 为新功能
# 或
git checkout -b fix/my-bugfix  # 为bug修复
```

### 3. 开发

1. 进行您的修改
2. 遵循[代码风格指南](DEVELOPER.md#代码风格指南)
3. 添加或修改测试，确保测试覆盖率不降低
4. 确保所有测试通过
   ```bash
   go test ./...
   ```

### 4. 提交

1. 确保您的代码通过了所有测试
2. 提交您的更改，使用清晰的提交消息：
   ```bash
   git add .
   git commit -m "feat: add new feature X" # 或 "fix: resolve issue with Y"
   ```
   我们建议遵循[Conventional Commits](https://www.conventionalcommits.org/)规范

3. 将您的分支推送到您的fork仓库：
   ```bash
   git push origin feature/my-feature
   ```

### 5. 创建Pull Request

1. 前往GitHub上您fork的仓库
2. 点击"Pull Request"按钮
3. 选择您刚刚推送的分支
4. 填写PR描述，包括：
   - 您解决了什么问题
   - 您是如何解决的
   - 可能需要审查者特别注意的地方
   - 相关的issue编号（如有）

### 6. 代码审查

1. 项目维护者将会审查您的PR
2. 根据反馈进行必要的修改
3. 一旦获得批准，您的贡献将被合并

## 开发指南

### 设置开发环境

1. 确保您已安装Go 1.20+
2. 安装开发依赖：
   ```bash
   go get -u github.com/stretchr/testify/assert
   go get -u github.com/spf13/cobra
   ```

### 测试

- 运行全部测试：`go test ./...`
- 检查测试覆盖率：`go test ./... -cover`
- 生成测试覆盖率报告：
  ```bash
  go test ./... -coverprofile=coverage.out
  go tool cover -html=coverage.out
  ```

### 编码规范

- 使用`gofmt`和`golint`确保代码质量
- 为所有导出的函数、类型和变量添加文档注释
- 错误处理必须明确
- 测试应涵盖正常和错误情况

## Pull Request指南

一个好的PR应该：

1. 聚焦于单一功能或bug修复
2. 包含适当的测试
3. 更新相关文档
4. 通过所有测试
5. 遵循代码风格指南

## 许可证

通过贡献，您同意您的贡献将在项目的[LICENSE](LICENSE)下发布。请注意，本项目商业使用需要获得授权。

## 问题与讨论

如有任何问题或需要讨论，请：

1. 在GitHub上创建一个Issue
2. 详细描述您的问题或想法
3. 添加相关的标签

感谢您的贡献！ 