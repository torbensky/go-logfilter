package example

import log "github.com/sirupsen/logrus"

func Foo() {
	log.Debug("File1")
	log.Info("File1")
	log.Warn("File1")
	log.Error("File1")
}
