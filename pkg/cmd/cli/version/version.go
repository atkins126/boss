package version

import (
	"fmt"

	"github.com/hashload/boss/internal/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print the Boss version",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
}

func printVersion() {
	v := version.Get()
	fmt.Println("Version       ", v.Version)
	fmt.Println("Git commit    ", v.GitCommit)
	fmt.Println("Git tree state", v.GitTreeState)
	fmt.Println("Go version    ", v.GoVersion)
}
