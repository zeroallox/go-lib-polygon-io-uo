package polyrest

import (
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

var json = jsoniter.ConfigDefault
var debugMode = false

func EnableDebug() {
	debugMode = true
	log.SetLevel(log.DebugLevel)
}
