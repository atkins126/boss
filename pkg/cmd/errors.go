package cmd

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		if err != context.Canceled {
			log.Errorf("An error occurred: %v\n", err)
		}
		os.Exit(1)
	}
}
