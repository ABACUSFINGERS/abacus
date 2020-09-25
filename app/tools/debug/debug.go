package debug

import (
	"log"
	"os"
)


func Log(format string, v ...interface{}) {
	if os.Getenv("DEBUG") == "1" {
		log.Printf(format, v...)
	}
}

func Error(format string, v ...interface{}) {
	log.Printf(format, v...)
}

