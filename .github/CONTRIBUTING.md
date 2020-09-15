## 贡献指南

你好！我们很高兴你想为这个项目做出贡献。 你的帮助对保持它的完美是至关重要的。

## 代码格式

  1. 函数名严格按照 PHP下划线 => Go大驼峰命名
     - array() => Array()
     - [array_change_key_case() => ArrayChangeKeyCase()](https://github.com/rhettli/php2go/blob/master/php/array_change_key_case.go)

  1. 缩进统一使用Tab
  1. 每次提交一个函数，注释以函数名开头
  1. 新增函数文件以PHP函数名命名单独提交，方便后期函数排序整理
  1. 不要提交二进制文件(本地测试可以使用go run命令)

示例:

```go

// FuncName - Description of function do
func FuncName(args Type) Type {
    return args
}

```

### 贡献代码

1. Fork [`https://github.com/rhettli/php2go`](https://github.com/rhettli/php2go) 并克隆到本地.
1. 创建新的分支：

    ```shell
    $ git checkout -b my-new-feature
    ```

1. 编写提交代码。

    ```shell
    $ git commit -am 'Add some feature'
    ```
    
1. Push 到你的分支。

    ```shell
    $ git push origin my-new-feature
    ```

1. 创建 Pull Request 并描述你完成的功能或者做出的修改。

## Resources

- [Contributing to Open Source on GitHub](https://guides.github.com/activities/contributing-to-open-source/)
- [Using Pull Requests](https://help.github.com/articles/using-pull-requests/)
- [GitHub Help](https://help.github.com)
