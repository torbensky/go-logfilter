package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func Foo(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Info("Foo is starting!")
	for i := 0; ; i++ {
		log.Debugf("Foo is doing thing %d", i)
		if i == 5 {
			log.Warn("Foo is 1/2 done!")
		}
		if i > 5 {
			log.Error("Uh-oh, I have a problem!")
			log.Debug("The problem is I lost my DB connection!")
		}
		time.Sleep(1 * time.Second)

		if i > 20 {
			break
		}
	}

	log.Info("Foo is done!")
}
