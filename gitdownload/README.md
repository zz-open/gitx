# gitdownload
- [github rest api](https://docs.github.com/en/rest/quickstart)
- [download-a-repository-archive-zip](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#download-a-repository-archive-zip)
- [download-a-repository-archive-tar](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#download-a-repository-archive-tar)
- [get-a-tree](https://docs.github.com/en/rest/git/trees?apiVersion=2022-11-28#get-a-tree)
- [get-repository-content](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#get-repository-content)

下载git仓库中的[指定文件]或者[目录]到[本地路径]

## 功能
### 下载单个文件
api 请求
```shell
GET /repos/{owner}/{repo}/contents/{path}

示例：https://api.github.com/repos/zzopen/mysqldoc/contents/cli/common.mk?ref=main
```

响应
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

下载方式：
- 通过响应信息中的 content | base64_decode
- 通过响应信息中的 download_url 直接下载
- 固定格式url转换下载 https://raw.githubusercontent.com/zzopen/mysqldoc/main/Makefile

### 下载指定目录
api 请求
```shell
GET /repos/{owner}/{repo}/contents/{path}

示例：https://api.github.com/repos/zzopen/mysqldoc/contents/src?ref=main
```

响应
```json
[
  {
    "name": ".gitignore",
    "path": "src/.gitignore",
    "sha": "0e41b127c228e4cddc40c303349ae55d42d1774d",
    "size": 10,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/.gitignore?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/blob/main/src/.gitignore",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/0e41b127c228e4cddc40c303349ae55d42d1774d",
    "download_url": "https://raw.githubusercontent.com/zzopen/mysqldoc/main/src/.gitignore",
    "type": "file",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/.gitignore?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/0e41b127c228e4cddc40c303349ae55d42d1774d",
      "html": "https://github.com/zzopen/mysqldoc/blob/main/src/.gitignore"
    }
  },
  {
    "name": "common",
    "path": "src/common",
    "sha": "b211e0cfd81b90493bd13c6e89047c0566610fea",
    "size": 0,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/common?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/tree/main/src/common",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/b211e0cfd81b90493bd13c6e89047c0566610fea",
    "download_url": null,
    "type": "dir",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/common?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/b211e0cfd81b90493bd13c6e89047c0566610fea",
      "html": "https://github.com/zzopen/mysqldoc/tree/main/src/common"
    }
  },
  {
    "name": "etc",
    "path": "src/etc",
    "sha": "0419fd733ee76ddf6f3b96105135fe428d7902ef",
    "size": 0,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/etc?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/tree/main/src/etc",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/0419fd733ee76ddf6f3b96105135fe428d7902ef",
    "download_url": null,
    "type": "dir",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/etc?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/0419fd733ee76ddf6f3b96105135fe428d7902ef",
      "html": "https://github.com/zzopen/mysqldoc/tree/main/src/etc"
    }
  },
  {
    "name": "internal",
    "path": "src/internal",
    "sha": "b7eb4c09ca61a2bb2239fdb9095f771ed1c87ba4",
    "size": 0,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/internal?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/tree/main/src/internal",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/b7eb4c09ca61a2bb2239fdb9095f771ed1c87ba4",
    "download_url": null,
    "type": "dir",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/internal?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/b7eb4c09ca61a2bb2239fdb9095f771ed1c87ba4",
      "html": "https://github.com/zzopen/mysqldoc/tree/main/src/internal"
    }
  },
  {
    "name": "mysqldoc.go",
    "path": "src/mysqldoc.go",
    "sha": "f149449f028ea5c7b253e60b4ae848517f4c28f0",
    "size": 99,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/mysqldoc.go?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/blob/main/src/mysqldoc.go",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/f149449f028ea5c7b253e60b4ae848517f4c28f0",
    "download_url": "https://raw.githubusercontent.com/zzopen/mysqldoc/main/src/mysqldoc.go",
    "type": "file",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/mysqldoc.go?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/f149449f028ea5c7b253e60b4ae848517f4c28f0",
      "html": "https://github.com/zzopen/mysqldoc/blob/main/src/mysqldoc.go"
    }
  },
  {
    "name": "ui",
    "path": "src/ui",
    "sha": "0bea968ac91c4d5f6e4732cae3f0205569399a84",
    "size": 0,
    "url": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/ui?ref=main",
    "html_url": "https://github.com/zzopen/mysqldoc/tree/main/src/ui",
    "git_url": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/0bea968ac91c4d5f6e4732cae3f0205569399a84",
    "download_url": null,
    "type": "dir",
    "_links": {
      "self": "https://api.github.com/repos/zzopen/mysqldoc/contents/src/ui?ref=main",
      "git": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/0bea968ac91c4d5f6e4732cae3f0205569399a84",
      "html": "https://github.com/zzopen/mysqldoc/tree/main/src/ui"
    }
  }
]
```

## github url 映射关系
```shell
# 递归列出全部文件
https://api.github.com/repos/zzopen/mysqldoc/git/trees/main?recursive=1

# 获取文件内容
https://raw.githubusercontent.com/zzopen/mysqldoc/main/Makefile

```
```shell
# 网页端查看目录
https://github.com/zzopen/mysqldoc/tree/main/test

# 网页端查看文件
https://github.com/zzopen/mysqldoc/blob/main/test/main.go
```