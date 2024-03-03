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
