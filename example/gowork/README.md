# Go.Work应用示例

> **go.work**是用于指定工作区中多个模块。

<br/>

# Work命令

**查看帮助：**
```shell
$ go help work
To determine whether the go command is operating in gowork mode, use
the "go env GOWORK" command. This will specify the gowork file being
used.

Usage:

        go work <command> [arguments]

The commands are:

        edit        edit gowork from tools or scripts
        init        initialize gowork file
        sync        sync gowork build list to modules
        use         add modules to gowork file

Use "go help work <command>" for more information about a command.
```
**初始化工作文件：**
```shell
$ cd gowork
$ go work init
```
**添加/升级模块：**
```shell
go work use projectA
go work use projectB
$ go work use
```

<br/>

# GoWork示例

- 创建两个Go项目`projectA`和`projectB`, 项目`projectA`调用`projectB`中的函数。工作目录如下：

```shell
$ tree ./gowork
./gowork
├── README.md
├── gowork
├── projectA
│   ├── go.mod
│   └── main.go
└── projectB
    ├── go.mod
    └── handler
        └── hello_handler.go        
```
- `projectB`实现函数实现

```go
// gowork/projectB/handler/hello_handler.go
package handler

import "fmt"

func Hello() {
	fmt.Println("This is an example of go.work")
}
```
- `projectA`调用`projectB`的函数
```go
// gowork/projectA/main.go
package main

import (
	"git.example.com/projectB/handler"
	"github.com/hollson/greet"
)

func main() {
	handler.Hello()
	greet.Hello()
}
```