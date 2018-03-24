package example

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	filter "github.com/torbensky/go-logfilter"
)

func Run() {
	config := `
		file1.go:debug,
		file2.go:warn,
		github.com/torbensky/go-logfilter:panic
	`

	// Set to debug so handlers are called at all levels to let our filters do a more granular log suppression.
	log.SetLevel(log.DebugLevel)
	aHook := ExampleHook{}

	f, err := filter.LoadConfig(config)
	if err != nil {
		panic(err) // demo!
	}

	// wrap the example hook with the hook filter
	filteredHook := filter.NewHookFilter(aHook, f)

	// register the wrapped, filtered hook with the log library
	log.AddHook(filteredHook)

	// Discard the default library output
	log.SetOutput(ioutil.Discard)

	// Call some stuff that produces logs
	Foo()
	Bar()
}
