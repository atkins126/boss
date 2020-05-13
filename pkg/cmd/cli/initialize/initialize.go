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
		},
	}
}
