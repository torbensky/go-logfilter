package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func Bar(wg sync.WaitGroup) {
	log.Info("Bar is starting!")
	for i := 0; i < 10; i++ {
		log.Debugf("Bar is doing thing %d", i)
		if i == 5 {
			log.Warn("Bar is 1/2 done!")
		}
		time.Sleep(1 * time.Second)
	}

	log.Info("Bar is done!")

	wg.Done()
}
