package text

import (
	log "github.com/sirupsen/logrus"
)

type Name struct {
	MyName string
}

func (name Name) name() string  {
	log.Infof("this is info")
	return "this is v1.0.0"

}

