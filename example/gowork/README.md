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
go work use projectC
$ go work use
```

<br/>

# GoWork示例

- 创建两个Go项目`projectA`和`projectB`, 项目`projectA`调用`projectB`中的函数。工作目录如下：

```shell
$ tree gowork/
gowork/
├── go.work
├── projectA
│   ├── go.mod
│   └── main.go
├── projectB
│   ├── go.mod		# 伪URL: git.example.com/projectB
│   └── handler.go
└── projectC
    ├── go.mod		# 伪Path: example/projectC
    └── handler.go
```
- `projectB`和`projectC`的`handler.go`实现函数实现:

```go
// gowork/projectB/handler.go
package projectB

import "fmt"

func Hello() {
	fmt.Println("Hello from projectB")
}
```

```go
// gowork/projectC/handler.go
package projectC

import "fmt"

func Hello() {
	fmt.Println("Hello from projectC")
}
```

- `projectA`调用`projectB`和`projectC`的函数

```js
// gowork/projectA/go.mods
module git.example.com/projectA

go 1.22.0

// 替换为相对路径
replace git.example.com/projectB => ../projectB

//伪UR必须添加requre
require (
	git.example.com/projectB v0.0.0-00010101000000-000000000000
	github.com/hollson/greet v1.0.0
)
```
- 编辑`main.go`调用函数：
```go
// gowork/projectA/main.go
package main

import (
	"example/projectC"              // 伪Path
	"git.example.com/user/projectB" // 伪URL(需要在go.mod添加requre项和热repleace配置)
	"github.com/hollson/greet"      // Github项目
)

func main() {
	greet.Hello()
	projectB.Hello()
	projectC.Hello()
}
```

**输出：**

```shell
$ go run main.go 
Hello World
Hello from projectB
Hello from projectC
```

