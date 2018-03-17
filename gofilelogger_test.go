package gofilelogger

import (
	"runtime"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLevel(t *testing.T) {
	lf := New()
	_, filename, _, _ := runtime.Caller(0)

	// Ensure the default level is working
	l := lf.GetFileLevel(filename)
	assert.Equal(t, l, logrus.InfoLevel)

	// Ensure the level can be overridden for a package
	lf.SetLevel(logrus.WarnLevel, "github.com/torbensky/gofilelogger")
	l = lf.GetFileLevel(filename)
	assert.Equal(t, logrus.WarnLevel, l, "package level")

	// Ensure the level can be overridden for a file
	lf.SetLevel(logrus.DebugLevel, "github.com/torbensky/gofilelogger/gofilelogger_test.go")
	l = lf.GetFileLevel(filename)
	assert.Equal(t, logrus.DebugLevel, l, "file level")

	// Try a mapping that doesn't use test filename
	lf.SetLevel(logrus.PanicLevel, "baz.go")
	l = lf.GetFileLevel("/foo/go/path/bar/baz/baz.go")
	assert.Equal(t, logrus.PanicLevel, l, "file level")

	// Try setting multiple paths at once
	paths := []string{"/unicorn/foo.go", "/dragon/bar.go", "/example/foo"}
	lf.SetLevels(logrus.DebugLevel, paths...)
	for _, p := range paths {
		l = lf.GetFileLevel(p)
		assert.Equal(t, logrus.DebugLevel, l, "multi-match")
	}
}

func TestCacheOff(t *testing.T) {
	lf := New()
	lf.UseCache = false

	lf.GetFileLevel("/foo/cache")
	assert.False(t, checkCache(lf.fileCache, "/foo/cache"))

	lf.SetLevel(logrus.WarnLevel, "/foo/cache")
	assert.False(t, checkCache(lf.fileCache, "/foo/cache"))

	l := lf.GetFileLevel("/foo/cache")
	assert.Equal(t, logrus.WarnLevel, l)
	assert.False(t, checkCache(lf.fileCache, "/foo/cache"))
}

func TestCacheOn(t *testing.T) {
	lf := New()

	assert.False(t, checkCache(lf.fileCache, "/foo/cache"))
	lf.GetFileLevel("/foo/cache")
	assert.True(t, checkCache(lf.fileCache, "/foo/cache"))

	lf.SetLevel(logrus.WarnLevel, "/foo/cache")
	assert.False(t, checkCache(lf.fileCache, "/foo/cache"))

	l := lf.GetFileLevel("/foo/cache")
	assert.Equal(t, logrus.WarnLevel, l)
	assert.True(t, checkCache(lf.fileCache, "/foo/cache"))
}

func checkCache(m *sync.Map, key string) bool {
	_, ok := m.Load(key)
	return ok
}
