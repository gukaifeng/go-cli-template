

# go-cli-template


这是一个 Golang 中命令行程序开发框架 `urfave/cli` 的一个开始模板。

你应该根据自己的需要增加或删除其中的内容。


## 1. 编译相关命令

```shell
make
```

会在此目录下编译出一个二进制文件，名为 `$PROGRAM`，在 Makefile 中配置。

```shell
make install
```

将编译出的二进制文件 `$PROGRAM` 移动到目录 `/usr/local/bin/` 下。


```shell
make clean
```

删除根目录下编译出的二进制文件 `$PROGRAM`。


## 2. 模板设计结构

```shell
$ tree .
.
├── autocomplete
│   ├── bash_autocomplete
│   ├── powershell_autocomplete.ps1
│   └── zsh_autocomplete
├── cmd
│   ├── cmd.go
│   ├── subcmd1.go
│   └── subcmd2.go
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md

2 directories, 11 files
```


根目录的 `main.go` 是入口，其内容十分简单，仅仅调用 `cmd/cmd.go:Main()` 以开始程序。

目录 `autocomplete/` 下是 `urfave/cli` 官方提供的三个在默认自动补全配置下分别适用 Bash、Zsh、PowerShell 的自动补全脚本。

`cmd/` 是此项目的核心目录，其中：

* `cmd/cmd.go` 是程序起点，也包含了命令行程序的相关配置，以及不含子命令时的相关操作方法。
* `cmd/subcmd1.go` 和 `cmd/subcmd2.go` 分别是两个子命令 `subcmd1` 和 `subcmd2` 的单独源文件，

在 `cmd/cmd.go` 中，我们把子命令独立出来了：
```go
app := &cli.App{
    // ...
    Commands: []*cli.Command{
        cmdSubcmd1(),
        cmdSubcmd2(),
    },
}
```

而在 `cmd/subcmd1.go` 和 `cmd/subcmd2.go` 中放置子命令的相关方法，以  `cmd/subcmd1.go` 为例：

```go
func cmdSubcmd1() *cli.Command {
	return &cli.Command{
        // ...
    }
}
```

这样能做到更好的项目层次结构。

