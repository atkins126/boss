package version

import (
	"github.com/hashload/boss/internal/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the Boss version",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
			// log.WithFields(log.Fields{
			// 	"animal": "walrus",
			// }).Info("A walrus appears")
		},
	}
}

func printVersion() {
	v := version.Get()
	log.Infoln("Version       ", v.Version)
	log.Infoln("Git commit    ", v.GitCommit)
	log.Infoln("Git tree state", v.GitTreeState)
	log.Infoln("Go version    ", v.GoVersion)
}
