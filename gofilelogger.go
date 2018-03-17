package gofilelogger

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var loggerConfig = sync.Map{} // thread safe, more performant than using regular map with mutex, I think
var fileCache = sync.Map{}
var UseCache = true

/*
	Sets a log level for a path
*/
func SetLevel(level logrus.Level, path string) {
	loggerConfig.Store(path, level)

	if !UseCache {
		return
	}

	// Bust the cache
	fileCache.Range(func(f, v interface{}) bool {
		if strings.Contains(f.(string), path) {
			fileCache.Delete(f)
		}
		return true
	})
}

/*
	Sets a log level for each path provided
*/
func SetLevels(level logrus.Level, paths ...string) {
	for _, p := range paths {
		SetLevel(level, p)
	}
}

/*
	Gets the "best" matching log level matching the given file. "best" in this case means the first file match it finds or any directory match.

	Defaults to InfoLevel when no match is found

	When multiple matches are possible, this will behave non-deterministically. This is based on the underlying map behavior.
*/
func GetFileLevel(file string) logrus.Level {
	// Check if we already know the level for this file
	cached, ok := fileCache.Load(file)
	if ok {
		return cached.(logrus.Level)
	}

	level := logrus.InfoLevel // default level if we don't find a match

	dir := filepath.Dir(file)
	loggerConfig.Range(func(k, v interface{}) bool {
		// If we match a file exactly, immediately return
		if strings.HasSuffix(file, k.(string)) {
			level = v.(logrus.Level)
			return false
		}

		// If we find a directory match, just remember it in case we don't find a file match
		if strings.HasSuffix(dir, k.(string)) {
			level = v.(logrus.Level)
		}

		return true
	})

	// Cache this result for faster lookup next time
	if UseCache {
		fileCache.Store(file, level)
	}

	return level
}
