# Gin脚手架工具 (GS)

[![许可证](https://img.shields.io/badge/license-Custom-blue.svg)](LICENSE)
[![Go版本](https://img.shields.io/badge/Go-1.20+-00ADD8.svg)](https://golang.org/)
[![测试覆盖率](https://img.shields.io/badge/coverage-90%25-brightgreen.svg)]()

**GS** 是一个强大的命令行工具，用于快速生成和搭建基于 [Gin框架](https://github.com/gin-gonic/gin) 的 Go Web 应用。它能帮助开发者轻松创建项目结构、控制器、模型、路由和服务等组件，大幅提高开发效率。

## 功能特点

- 🚀 快速初始化标准结构的Gin项目
- 📁 自动生成符合最佳实践的项目文件结构
- 🧩 轻松创建控制器、模型、路由和服务组件
- 🔄 一键生成完整的CRUD资源
- ⚙️ 可定制的模板系统
- 🧪 完善的测试覆盖

## 安装

### 通过go install安装

```bash
go install github.com/yggai/gs/cmd/gs@latest
```

### 从源码构建

```bash
git clone https://github.com/yggai/gs.git
cd gs
go build -o gs ./cmd/gs
```

将生成的可执行文件移动到`$PATH`中的目录，即可在任何位置使用`gs`命令。

## 快速开始

### 创建新项目

```bash
# 创建名为myapp的新项目
gs init myapp

# 指定Go模块名称
gs init myapp --module github.com/username/myapp
```

### 生成组件

```bash
# 创建控制器
gs create controller User

# 创建模型
gs create model User

# 创建路由
gs create router User

# 创建服务
gs create service User

# 一次性创建完整资源（控制器、模型、路由和服务）
gs create resource User
```

## 项目结构

使用`gs`初始化的项目结构如下：

```
myapp/
├── config/             # 配置文件
├── controllers/        # 控制器
├── models/             # 数据模型
├── services/           # 业务逻辑层
├── routers/            # 路由定义
├── middleware/         # 中间件
├── utils/              # 工具函数
├── go.mod              # Go模块定义
├── main.go             # 应用入口
└── README.md           # 项目说明
```

## 自定义模板

GS使用可自定义的模板系统，您可以根据自己的需求修改模板。默认模板位于`templates`目录中：

```
templates/
├── project/            # 项目模板
└── component/          # 组件模板
    ├── controller/     # 控制器模板
    ├── model/          # 模型模板
    ├── router/         # 路由模板
    └── service/        # 服务模板
```

若要使用自定义模板，请设置环境变量：

```bash
export GS_TEMPLATES_DIR=/path/to/your/templates
```

## 命令参考

### 全局标志

- `--version`, `-v` - 显示版本信息

### init 命令

初始化新的Gin项目。

```bash
gs init [项目名称] [flags]
```

**标志:**

- `--module`, `-m` - 指定Go模块名称 (默认为项目名称)
- `--force`, `-f` - 强制初始化，即使目标目录已存在

### create 命令

创建各种组件。

```bash
gs create [组件类型] [名称] [flags]
```

**组件类型:**

- `controller` - 创建控制器
- `model` - 创建模型
- `router` - 创建路由
- `service` - 创建服务
- `resource` - 创建完整资源（包含上述所有组件）

**标志:**

- `--force`, `-f` - 强制创建，覆盖已存在的文件

## 开发

### 先决条件

- Go 1.20+

### 构建与测试

```bash
# 构建
go build -o gs ./cmd/gs

# 运行测试
go test ./...

# 查看测试覆盖率
go test ./... -cover
```

## 许可证

本项目采用自定义开源许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

**注意:** 商业使用需获得授权。

## 贡献指南

欢迎提交问题报告、功能请求和代码贡献。请遵循以下步骤：

1. Fork本项目
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的修改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开Pull Request

## 联系方式

如有问题或需要商业授权，请联系：[您的联系信息]
