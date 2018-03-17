package gofilelogger

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func LoadConfig(config string) (*LogFilter, error) {
	lf := New()

	// Each log level mapping is separated by comma
	for _, s := range strings.Split(config, ",") {

		// Each file path is separated from log level by a ":"
		pathAndLevel := strings.Split(strings.TrimSpace(s), ":")
		if len(pathAndLevel) != 2 {
			return nil, errors.Errorf("unexpected syntax in filter: %s", s)
		}

		l, err := logrus.ParseLevel(pathAndLevel[1])
		if err != nil {
			return nil, errors.Wrap(err, "unable to parse log level")
		}

		lf.SetLevel(l, pathAndLevel[0])
	}

	return lf, nil
}
