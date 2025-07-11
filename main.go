package main

import (
	"os"
	
	"github.com/yggai/gs/cmd/gs"
)

func main() {
	// 执行根命令
	if err := gs.Execute(); err != nil {
		os.Exit(1)
	}
} 