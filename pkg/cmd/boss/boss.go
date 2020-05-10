package boss

import (
	"github.com/hashload/boss/pkg/cmd/cli/version"
	"github.com/spf13/cobra"
)

func NewBossCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "boss",
		Short: "Dependency Manager for Delphi",
		Long:  "Dependency Manager for Delphi",
	}

	root.AddCommand(version.NewVersionCommand())

	return root
}
