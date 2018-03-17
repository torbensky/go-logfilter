package example

import log "github.com/sirupsen/logrus"

func Bar() {
	log.Debug("File2")
	log.Info("File2")
	log.Warn("File2")
	log.Error("File2")
}
