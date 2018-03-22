package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	filter "github.com/torbensky/gofilelogger"
	"github.com/torbensky/gofilelogger/example"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	// Logging config
	config := os.Getenv("LOG_LEVELS")
	fmt.Println("Configuration:")
	fmt.Println(config)

	// Set to debug so handlers are called at all levels to let our filters do a more granular log suppression.
	log.SetLevel(log.DebugLevel)
	aHook := example.ExampleHook{}

	f, err := filter.LoadConfig(config)
	if err != nil {
		panic(err) // demo!
	}
	fmt.Println(f)

	// wrap the example hook with the hook filter
	filteredHook := filter.NewHookFilter(aHook, f)

	// register the wrapped, filtered hook with the log library
	log.AddHook(filteredHook)

	// Discard the default library output
	log.SetOutput(ioutil.Discard)

	// Run some async processes to simulate different system modules doing their thang
	var wg sync.WaitGroup
	wg.Add(3)

	go Foo(wg)
	go Bar(wg)
	go Out(wg)

	wg.Wait()
}
