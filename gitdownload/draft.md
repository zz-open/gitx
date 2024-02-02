### 下载指定目录
通过以下github api 递归遍历目录下载
```shell
GET /repos/{owner}/{repo}/git/trees/{tree_sha}

- tree_sha 只能是分支名称或者目录的SHA1值，不能是文件的SHA1值
- ?recursive=1 递归获取文件

例如：https://api.github.com/repos/zzopen/mysqldoc/git/trees/798b5288780a6771df94932293a81fdc47eeb65b
```

响应
```json
{
  "sha": "798b5288780a6771df94932293a81fdc47eeb65b",
  "url": "https://api.github.com/repos/zzopen/mysqldoc/git/trees/798b5288780a6771df94932293a81fdc47eeb65b",
  "tree": [
    {
      "path": "common.mk",
      "mode": "100644",
      "type": "blob",
      "sha": "c7c07afb82eec8edf5b385644be4e56f19a8f8c3",
      "size": 613,
      "url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/c7c07afb82eec8edf5b385644be4e56f19a8f8c3"
    },
    {
      "path": "server.mk",
      "mode": "100644",
      "type": "blob",
      "sha": "3ad5eb193f34177d7c50a861c52398eec67c15c2",
      "size": 1951,
      "url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/3ad5eb193f34177d7c50a861c52398eec67c15c2"
    },
    {
      "path": "ui.mk",
      "mode": "100644",
      "type": "blob",
      "sha": "4b66b54b92f73af7c934beb244c94173e1f7c8d6",
      "size": 393,
      "url": "https://api.github.com/repos/zzopen/mysqldoc/git/blobs/4b66b54b92f73af7c934beb244c94173e1f7c8d6"
    }
  ],
  "truncated": false
}
```