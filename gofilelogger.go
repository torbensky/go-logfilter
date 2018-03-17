package gofilelogger

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type LogFilter struct {
	config    *sync.Map // thread safe map, from reading sounds like this should be more efficient than manually managing a mutex
	fileCache *sync.Map // thread safe map, from reading sounds like this should be more efficient than manually managing a mutex
	UseCache  bool      // whether or not to use a cache of filenames seen and the log levels found
}

func NewLogFilter() *LogFilter {
	return &LogFilter{
		config:    &sync.Map{},
		fileCache: &sync.Map{},
		UseCache:  true,
	}
}

/*
	Sets a log level for a path
*/
func (lf *LogFilter) SetLevel(level logrus.Level, path string) {
	lf.config.Store(path, level)

	if !lf.UseCache {
		return
	}

	// Bust the cache
	lf.fileCache.Range(func(f, v interface{}) bool {
		// A simple and aggressive prune, eliminating any possible match. Should at least be a superset of items affected by the new level setting.
		if strings.Contains(f.(string), path) {
			lf.fileCache.Delete(f)
		}
		return true
	})
}

/*
	Sets a log level for each path provided
*/
func (lf *LogFilter) SetLevels(level logrus.Level, paths ...string) {
	for _, p := range paths {
		lf.SetLevel(level, p)
	}
}

/*
	Gets the "best" matching log level matching the given file. "best" in this case means the first file match it finds or any directory match.

	Defaults to InfoLevel when no match is found

	When multiple matches are possible, this will behave non-deterministically. This is based on the underlying map behavior.
*/
func (lf *LogFilter) GetFileLevel(file string) logrus.Level {
	// Check if we already know the level for this file
	cached, ok := lf.fileCache.Load(file)
	if ok {
		return cached.(logrus.Level)
	}

	level := logrus.InfoLevel // default level if we don't find a match

	dir := filepath.Dir(file)
	lf.config.Range(func(k, v interface{}) bool {
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
	if lf.UseCache {
		lf.fileCache.Store(file, level)
	}

	return level
}
