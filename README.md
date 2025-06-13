# Template

## 架构

```shell
├─cmd           # 启动文件
├─config        # 配置文件
├─internal
│  ├─data       # Data 层
│  ├─handler    # Handler 层
│  └─server     # Server 层
├─logs          # 运行时日志
├─tmp           # 应用构建
└─utils
```

## 构建说明

### 版本注入

在构建应用时，我们可以通过 `-ldflags` 参数注入版本信息：

```shell
go build -ldflags "-X main.Version=x.y.z"
```

这个命令的作用是：

- `-ldflags`: 传递给链接器的参数
- `-X`: 用于设置包中变量的值
- `main.Version`: 指定要设置的变量（包名.变量名）
- `x.y.z`: 要注入的版本号

这样在运行时就可以通过 `main.Version` 获取到构建时注入的版本号。

## 安装

```shell
# Task 工具：相当于 Makefile
winget install Task.Task
```

## 运行

```shell
# 生成代码
task generate

# 构建应用
task build

# 运行实例
task run
```
