package replicator

import (
  log "github.com/sirupsen/logrus"
)

func replicate() string {
	return "replicate!!"
}

func Start() {
	log.Info(replicate())
}
