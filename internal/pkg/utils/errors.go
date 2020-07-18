package utils

import (
	"context"
	"log"
)

func CheckError(err error) {
	if err != nil {
		if err != context.Canceled {
			log.Fatalf("An error occurred: %v\n", err)
		}
	}
}
