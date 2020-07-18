package initialize

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Long:  `This command initialize a new project`,
		Run: func(cmd *cobra.Command, args []string) {
			runInitialize()
		},
	}
}

func runInitialize() {
	printHead()
}

func printHead() {
	println(`
This utility will walk you through creating a boss.json file.
It only covers the most common items, and tries to guess sensible defaults.
		 
Use 'boss install <pkg>' afterwards to install a package and
save it as a dependency in the boss.json file.
Press ^C at any time to quit.`)
}
