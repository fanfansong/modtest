package simpletest

import "github.com/apex/log"

func String() {
	log.WithFields(log.Fields{
		"package":  "simpletest",
		"function": "String()",
	}).Info("info")
}
