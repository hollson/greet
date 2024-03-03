module git.example.com/user/projectA

go 1.22.0

// 替换为相对路径
replace git.example.com/user/projectB => ../projectB

require (
	git.example.com/user/projectB v0.0.0-00010101000000-000000000000
	github.com/hollson/greet v1.0.0
)
