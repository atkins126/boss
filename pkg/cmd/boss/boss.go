package boss

import (
	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/pkg/cmd/cli/completion"
	"github.com/hashload/boss/pkg/cmd/cli/initialize"
	"github.com/hashload/boss/pkg/cmd/cli/login"
	"github.com/hashload/boss/pkg/cmd/cli/upgrade"
	"github.com/hashload/boss/pkg/cmd/cli/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewBossCommand(name string) *cobra.Command {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	config := configuration.LoadConfiguration()

	root := &cobra.Command{
		Use:   name,
		Short: "Dependency Manager for Delphi",
		Long:  "Dependency Manager for Delphi",
	}

	config.BindFlags(root)

	root.AddCommand(version.NewCmdVersion())
	root.AddCommand(initialize.NewCmdInitialize())
	root.AddCommand(login.NewCmdLogin(config))
	root.AddCommand(upgrade.NewCmdUpgrade())
	root.AddCommand(completion.NewCmdCompletion())

	return root
}
