package gofilelogger

import (
	"runtime"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLevel(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)

	// Ensure the default level is working
	l := GetFileLevel(filename)
	assert.Equal(t, l, logrus.InfoLevel)

	// Ensure the level can be overridden for a package
	SetLevel(logrus.WarnLevel, "github.com/torbensky/gofilelogger")
	l = GetFileLevel(filename)
	assert.Equal(t, logrus.WarnLevel, l, "package level")

	// Ensure the level can be overridden for a file
	SetLevel(logrus.DebugLevel, "github.com/torbensky/gofilelogger/gofilelogger_test.go")
	l = GetFileLevel(filename)
	assert.Equal(t, logrus.DebugLevel, l, "file level")

	// Try a mapping that doesn't use test filename
	SetLevel(logrus.PanicLevel, "baz.go")
	l = GetFileLevel("/foo/go/path/bar/baz/baz.go")
	assert.Equal(t, logrus.PanicLevel, l, "file level")

	// Try setting multiple paths at once
	paths := []string{"/unicorn/foo.go", "/dragon/bar.go", "/example/foo"}
	SetLevels(logrus.DebugLevel, paths...)
	for _, p := range paths {
		l = GetFileLevel(p)
		assert.Equal(t, logrus.DebugLevel, l, "multi-match")
	}
}
