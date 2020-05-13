package main

import (
	"os"
	"path/filepath"

	"github.com/hashload/boss/pkg/cmd"
	"github.com/hashload/boss/pkg/cmd/boss"
)

func main() {
	baseName := filepath.Base(os.Args[0])

	err := boss.NewBossCommand(baseName).Execute()
	cmd.CheckError(err)
}
