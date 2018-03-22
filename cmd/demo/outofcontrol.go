package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func Out(wg sync.WaitGroup) {
	for {
		log.Error("I just like to be noticed, but I'm not important :P")
		time.Sleep(200 * time.Millisecond)
	}
	wg.Done()
}
