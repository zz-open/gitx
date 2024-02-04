# download
下载github仓库资源到本地

- 支持下载单个文件
- 支持下载仓库内的子目录或子文件夹
- 支持下载整个仓库文件(.zip或者.tar.gz)

## 项目用到的 github api
- [github rest api](https://docs.github.com/en/rest/quickstart)
- [download-a-repository-archive-zip](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#download-a-repository-archive-zip)
- [download-a-repository-archive-tar](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#download-a-repository-archive-tar)
- [get-repository-content](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#get-repository-content)
- [get-a-tree](https://docs.github.com/en/rest/git/trees?apiVersion=2022-11-28#get-a-tree)
- [get-a-blob](https://docs.github.com/en/rest/git/blobs?apiVersion=2022-11-28#get-a-blob)
- https://raw.githubusercontent.com/{username}/{repo}/{branch}/path

## 原理
### 下载单个文件

方案:
- 使用固定格式url下载,例如: https://raw.githubusercontent.com/zzopen/mysqldoc/main/Makefile
- 通过contents api响应信息中的 content 字段，base64_decode 即可 
- 通过contents api响应信息中的 download_url(https://raw.githubusercontent.com/xxx) 直接下载

api 请求:
```shell
GET /repos/{owner}/{repo}/contents/{path}

示例：https://api.github.com/repos/zzopen/mysqldoc/contents/cli/common.mk?ref=main
```

响应:
```json
{
  "name": "common.mk",
  "path": "cli/common.mk",
  "sha": "c7c07afb82eec8edf5b385644be4e56f19a8f8c3",
  "size": 613,
  "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/cli/common.mk?ref=main",
  "html_url": "https://github.com/zzopen/mysqldoc/blob/main/cli/common.mk",
  "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/c7c07afb82eec8edf5b385644be4e56f19a8f8c3",
  "download_url": "https://raw.githubusercontent.com/zzopen/mysqldoc/main/cli/common.mk",
  "type": "file",
  "content": "IyBjb21tYW5kCkdPPWdvCkdPX0NNRD1HTzExMU1PRFVMRT1vbgpHT19WRVQ9\nJChHTykgdmV0CkdPX0JVSUxEPSQoR08pIGJ1aWxkCkdPX0lNUE9SVFM9Z29p\nbXBvcnRzCkdPX0ZNVD1nb2ZtdApHT19DVEw9Z29jdGwKCiMgcGF0aApNQUtF\nRklMRV9QQVRIPSQoYWJzcGF0aCAkKGxhc3R3b3JkICQoTUFLRUZJTEVfTElT\nVCkpKQpNQUtFRklMRV9ESVI9JChhYnNwYXRoICQoZGlyICQoTUFLRUZJTEVf\nUEFUSCkpKQpDVVJSRU5UX1BBVEg9JChNQUtFRklMRV9ESVIpClJPT1RfUEFU\nSD0kKGFic3BhdGggJChDVVJSRU5UX1BBVEgpLy4uLykKUFJPSkVDVF9ST09U\nX1BBVEg9JChST09UX1BBVEgpCiMg5pyN5Yqh56uv5qC555uu5b2V57ud5a+5\n6Lev5b6EClNSQ19QQVRIPSQoUFJPSkVDVF9ST09UX1BBVEgpL3NyYwojIOac\njeWKoeerr+WtmOaUvuWJjeerr+aWh+S7tuagueebruW9lee7neWvuei3r+W+\nhApHT19VSV9QQVRIPSQoU1JDX1BBVEgpL3VpCiMg5YmN56uv5qC555uu5b2V\n57ud5a+56Lev5b6EClVJX1BBVEg9JChQUk9KRUNUX1JPT1RfUEFUSCkvdWkK\nIyDmnI3liqHnq6/miZPljIXovpPlh7rot6/lvoQKT1VUX1BBVEg9JChQUk9K\nRUNUX1JPT1RfUEFUSCkvbXlzcWxkb2Mtb3V0Cg==\n",
  "encoding": "base64",
  "_links": {
    "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/cli/common.mk?ref=main",
    "git": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/c7c07afb82eec8edf5b385644be4e56f19a8f8c3",
    "html": "https://github.com/zzopen/mysqldoc/blob/main/cli/common.mk"
  }
}
```

### 下载子目录
分为以下几个步骤

#### fetch contents
获取指定路径下的一级目录路径

api 请求:
```shell
GET /repos/{owner}/{repo}/contents/{path}

示例：https://api.github.com/repos/zzopen/mysqldoc/contents/cli?ref=main

此api支持的文件数量较少，所以适合获取一级目录
```

响应:
```json
[
  {
    "name": "common.mk",
    "path": "cli/common.mk",
    "sha": "c7c07afb82eec8edf5b385644be4e56f19a8f8c3",
    "size": 613,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/cli/common.mk?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/blob/main/cli/common.mk",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/c7c07afb82eec8edf5b385644be4e56f19a8f8c3",
    "download_url": "https://raw.githubusercontent.com/zzopen/mysqldoc/main/cli/common.mk",
    "type": "file",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/cli/common.mk?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/c7c07afb82eec8edf5b385644be4e56f19a8f8c3",
      "html": "https://github.com/zzopen/mysqldoc/blob/main/cli/common.mk"
    }
  },
  {
    "name": "server.mk",
    "path": "cli/server.mk",
    "sha": "3ad5eb193f34177d7c50a861c52398eec67c15c2",
    "size": 1951,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/cli/server.mk?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/blob/main/cli/server.mk",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/3ad5eb193f34177d7c50a861c52398eec67c15c2",
    "download_url": "https://raw.githubusercontent.com/zzopen/mysqldoc/main/cli/server.mk",
    "type": "file",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/cli/server.mk?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/3ad5eb193f34177d7c50a861c52398eec67c15c2",
      "html": "https://github.com/zzopen/mysqldoc/blob/main/cli/server.mk"
    }
  }
]
```

#### fetch trees
递归所有的子目录下的一级目录
**注意：此api支持的文件数量较多，所以可通过增加递归参数获取所有文件**

api 请求:
```shell
GET /repos/{owner}/{repo}/git/trees/{tree_sha}?recursive=1

- tree_sha 只能是分支名称或者目录的SHA1值，不能是文件的SHA1值
- ?recursive=1 递归获取文件

示例：https://api.github.com/repos/zzopen/mysqldoc/git/trees/b211e0cfd81b90493bd13c6e89047c0566610fea?recursive=1
```

响应:
```json
{
  "sha": "b211e0cfd81b90493bd13c6e89047c0566610fea",
  "url": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/b211e0cfd81b90493bd13c6e89047c0566610fea",
  "tree": [
    {
      "path": "embed",
      "mode": "040000",
      "type": "tree",
      "sha": "6f669ade2b64aa1a6221d37072ea6ff236c292c9",
      "url": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/6f669ade2b64aa1a6221d37072ea6ff236c292c9"
    },
    {
      "path": "embed/config.yaml.tpl",
      "mode": "100644",
      "type": "blob",
      "sha": "9cb2662e1d26537077cd55af5d8a4331f4f7fcf1",
      "size": 141,
      "url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/9cb2662e1d26537077cd55af5d8a4331f4f7fcf1"
    },
  ],
  "truncated": false
}
```
#### fetch blobs
并发下载文件

api 请求
```shell
GET /repos/{owner}/{repo}/git/blobs/{tree_sha}

示例：https://api.github.com/repos/zzopen/mysqldoc/git/blobs/9cb2662e1d26537077cd55af5d8a4331f4f7fcf1
```

响应
```json
{
  "sha": "9cb2662e1d26537077cd55af5d8a4331f4f7fcf1",
  "node_id": "B_kwDOKf5MF9oAKDljYjI2NjJlMWQyNjUzNzA3N2NkNTVhZjVkOGE0MzMxZjRmN2ZjZjE",
  "size": 141,
  "url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/9cb2662e1d26537077cd55af5d8a4331f4f7fcf1",
  "content": "UG9ydDogNzY1NApBdXRvT3BlbkRlZmF1bHRCcm93c2VyOiB0cnVlCkNyZWF0\nZVNxbEZpbGU6IHRydWUKTXlzcWw6CiAgSG9zdDogMTI3LjAuMC4xCiAgUG9y\ndDogMzMwNgogIFVzZXJuYW1lOiByb290CiAgUGFzc3dvcmQ6CiAgRGJOYW1l\nOiB0ZXN0\n",
  "encoding": "base64"
}

content | base_decode 之后的结果保存到文件中即可
```

### 下载整个项目
方案:
- https://api.github.com/repos/zzopen/mysqldoc/zipball/main (api)
- https://api.github.com/repos/zzopen/mysqldoc/tarball/main (api)
- https://github.com/zz-guide/go-guide/archive/main.zip (api)
- https://github.com/zz-guide/go-guide/archive/refs/heads/main.zip (UI界面显示)

重点观察response header中的以下字段:
```shell
content-disposition: attachment; filename=go-guide-main.zip
Content-Type: application/zip
```
响应数据是二进制的，根据Content-Type创建对应后缀的文件保存数据即可

api 请求:
```shell
GET /repos/{owner}/{repo}/tarball/{ref}

示例：https://api.github.com/repos/zzopen/mysqldoc/tarball/main
```

响应:
```json
Status: 302
```

api 请求
```shell
GET /repos/{owner}/{repo}/zipball/{ref}

示例：https://api.github.com/repos/zzopen/mysqldoc/zipball/main
```

响应
```json
Status: 302
```

curl 请求:
```shell
curl -L \
-H "Accept: application/vnd.github+json" \
-H "X-GitHub-Api-Version: 2022-11-28" \
https://api.github.com/repos/zzopen/mysqldoc/zipball/main \
--output ./mysqldoc.zip
```

## 依赖的第三方库
```shell
# 单元测试
go get -u github.com/stretchr/testify

# 协程池，文件数量较多的时候，不能无限开协程，需要使用池化技术控制数量
go get -u github.com/panjf2000/ants/v2
```