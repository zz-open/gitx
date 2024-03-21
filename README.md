# zb
命令行工具，内置了一些常用的功能，例如：下载github仓库资源等

## 依赖第三方库
```go
go get -u github.com/spf13/cobra@latest

go install github.com/spf13/cobra-cli@latest
```

## 如何使用
- go install github.com/zz-open/zb@latest
- 下载release文件进行安装

## 命令
| 命令 | 说明 |
| --- | --- |
| zb ghd | 下载github项目指定文件 |
| zb dns | 输出数据库dsn示例  |