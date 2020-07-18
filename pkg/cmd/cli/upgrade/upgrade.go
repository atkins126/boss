package upgrade

import (
	"github.com/hashload/boss/internal/pkg/boss"
	"github.com/spf13/cobra"
)

func NewUpgradeCommand() *cobra.Command {
	var preRelease bool
	upgradeCmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade a cli",
		Long:  `This command upgrade the client version`,
		Run: func(cmd *cobra.Command, args []string) {
			boss.Upgrade(preRelease)
		},
	}

	upgradeCmd.Flags().BoolVar(&preRelease, "dev", false, "Pre-release")

	return upgradeCmd
}
