package example

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type ExampleHook struct{}

func (h ExampleHook) Fire(entry *log.Entry) error {
	fmt.Printf("%s:%s\n", entry.Level, entry.Message)
	return nil
}

func (h ExampleHook) Levels() []log.Level {
	return log.AllLevels
}
