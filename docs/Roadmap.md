# Roadmap
- ghd
- dsn
- dd机器人
- ssh上传和下载
- 使用默认浏览器打开网址
- 下载m3u8视频


# 任务清单
- [ ] 内置通用钉钉机器人发送消息(因为api升级的缘故，先搁置)
- [ ] 调研是否能打开浏览器某个页面进行操作
- [ ] 内置通用服务器ssh上传和下载文件或者目录
- [ ] 内置下载m3u8视频功能
- [ ] 内置生成ascii art text
- [ ] 内置使用默认浏览器打开指定url
- [ ] 内置打开常用文档地址
- [ ] 输出一个dockerfile
- [ ] 输出一个k8s配置文件
- [x] 找到了一个支持prompt的库
- [x] 分析apache日志和nginx日志

- [ ] 考虑把密码本功能做到这个工具里，全套的，相对麻烦
- [ ] 内置歌曲查询api，内置页面可以在线播放，还可以内置歌单
- [ ] 集成chatgpt


# 安装
```shell
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest

cobra-cli init

go get -u github.com/jmoiron/sqlx
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/cobra@latest
go get -u github.com/spf13/viper
go get -u github.com/gin-contrib/cors
go get -u github.com/gookit/color
go get -u github.com/duke-git/lancet/v2
go get -u github.com/alibabacloud-go/dingtalk/
```

# 参考网址
- [playwright-go](https://playwright-community.github.io/playwright-go/)
- [playwright-go](https://github.com/playwright-community/playwright-go)
- [go-rod](https://go-rod.github.io/#/selectors/README)
- [go-rod](https://github.com/go-rod/rod)
- [robotgo](https://github.com/go-vgo/robotgo)
- [automaxprocs](https://github.com/uber-go/automaxprocs)
- [goreleaser](https://github.com/goreleaser/goreleaser)
- [gophish](https://getgophish.com/)
- [go-openai](https://github.com/sashabaranov/go-openai)
- [syncthing](https://github.com/syncthing/syncthing)
- [gotty](https://github.com/yudai/gotty)


# FAQ