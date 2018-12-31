package main

import (
  log "github.com/sirupsen/logrus"
)

func helloworld() string {
	return "Hello World!!"
}

func main() {
	log.Info(helloworld())
}
