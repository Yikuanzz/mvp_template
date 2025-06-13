# MVP Template 项目说明

## 项目架构

```
├─cmd           # 启动入口（main.go、wire.go等）
├─config        # 配置文件（config.toml等）
├─internal
│  ├─data       # Data 层：为 Handler 层提供数据接口实现
│  ├─handler    # Handler 层：业务逻辑处理，面向 Server 层 API
│  └─server     # Server 层：对外接口定义，面向前端
├─logs          # 运行时日志
├─tmp           # 应用构建产物
└─utils         # 工具包
```

## 配置说明

主配置文件：`config/config.toml`  
示例内容：

```toml
[auth]
secret_key = "whzhsk123456."
expires = 24

[server]
host = "localhost"
port = 8080
mode = "debug"
read_timeout = 60
write_timeout = 60
idle_timeout = 120

[mysql]
host = "localhost"
port = "3306"
username = "root"
password = "123456"
database = "mvp_db"
```

## 构建与运行

本项目推荐使用 [Task](https://taskfile.dev) 工具（类似 Makefile，支持 Windows！）

### 安装 Task

```shell
winget install Task.Task
```

### 常用命令

```shell
# 生成依赖代码（如 wire 依赖注入）
task generate

# 构建应用
task build

# 运行实例
task run
```

## 版本注入

构建时可通过 `-ldflags` 注入版本号：

```shell
go build -ldflags "-X main.Version=x.y.z"
```

## 约定式提交（Commitlint + Husky）

本项目已集成 commitlint + husky，自动校验 commit message，兼容 Windows！

- 钩子文件：`.husky/commit-msg`
- 校验规则：`commitlint.config.js`

### 支持的提交类型（type）

```text
fix      # 修复 bug
feat     # 新功能
docs     # 文档变更
style    # 代码格式（不影响功能，如空格、分号等）
refactor # 代码重构（既不是新增功能，也不是修 bug）
perf     # 性能优化
test     # 增加测试
revert   # 回滚提交
chore    # 构建过程或辅助工具变动
build    # 构建相关变更
ci       # 持续集成相关变更
```

### 提交信息格式

```text
<type>(<scope>): <subject>
# 空一行
<body>
# 空一行
<footer>
```

- `<type>`：提交类型，见上表
- `<scope>`：影响范围（可选）
- `<subject>`：简要描述
- `<body>`：详细描述（可选）
- `<footer>`：关联 issue 或 BREAKING CHANGE（可选）

### 示例

```text
feat(user): 新增用户注册接口

实现了用户注册的基本流程，包含邮箱校验和密码加密。

Closes #12
```

```text
fix: 修复登录时的密码校验 bug
```


## 目录分层说明

- **cmd/**：应用启动入口
- **config/**：配置文件及解析
- **internal/data/**：数据层，面向 Handler 层的接口实现
- **internal/handler/**：业务逻辑层，面向 Server 层 API
- **internal/server/**：服务层，接口定义
- **utils/**：通用工具包

