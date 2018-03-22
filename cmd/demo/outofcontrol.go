package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func Out(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		log.Error("I just like to be noticed, but I'm not important :P")
		time.Sleep(200 * time.Millisecond)

		if i > 100 {
			break
		}
	}
}
