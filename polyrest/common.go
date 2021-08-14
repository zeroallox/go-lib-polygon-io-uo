package polyrest

import (
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"time"
)

var json = jsoniter.ConfigDefault
var debugMode = false
var retryOn504 = false
var maxRetryCount = 50
var retryInterval = time.Second * 5

func EnableDebug() {
	debugMode = true
	log.SetLevel(log.DebugLevel)
}

func EnableAutoRetry() {
	retryOn504 = true
}
