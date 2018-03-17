package gofilelogger

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var loggerConfig = sync.Map{} // thread safe, more performant than using regular map with mutex, I think

/*
	Sets a log level for a path
*/
func SetLevel(level logrus.Level, path string) {
	loggerConfig.Store(path, level)
}

/*
	Sets a log level for each path provided
*/
func SetLevels(level logrus.Level, paths ...string) {
	for _, p := range paths {
		loggerConfig.Store(p, level)
	}
}

/*
	Gets the "best" matching log level matching the given file. "best" in this case means the first file match it finds or any directory match.

	Defaults to InfoLevel when no match is found

	When multiple matches are possible, this will behave non-deterministically. This is based on the underlying map behavior.
*/
func GetFileLevel(file string) (logrus.Level, bool) {
	dir := filepath.Dir(file)
	level := logrus.InfoLevel // default level if we don't find a match
	found := false
	loggerConfig.Range(func(k, v interface{}) bool {
		// If we match a file exactly, immediately return
		if strings.HasSuffix(file, k.(string)) {
			level = v.(logrus.Level)
			found = true
			return false
		}

		// If we find a directory match, just remember it in case we don't find a file match
		if strings.HasSuffix(dir, k.(string)) {
			level = v.(logrus.Level)
			found = true
		}

		return true
	})

	return level, found
}
