package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"things": "aaa",
	}).Infof("Is this working")
}
