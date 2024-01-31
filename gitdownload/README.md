# gitdownload
- [download-a-repository-archive-zip](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#download-a-repository-archive-zip)
- [download-a-repository-archive-tar](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#download-a-repository-archive-tar)

下载git仓库中的指定文件或者文件夹

## 构思
### repo url解析
根据给定url解析信息

## 注意
```shell
# 递归列出全部文件
https://api.github.com/repos/zzopen/mysqldoc/git/trees/main?recursive=1

# 获取文件内容
https://raw.githubusercontent.com/zzopen/mysqldoc/main/Makefile

# 查看目录
https://github.com/zzopen/mysqldoc/tree/main/test

# 查看文件
https://github.com/zzopen/mysqldoc/blob/main/test/main.go
```