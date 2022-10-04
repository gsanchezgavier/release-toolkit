package main

import (
	"os"

	"github.com/gsanchezgavier/release-toolkit/src/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := app.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
