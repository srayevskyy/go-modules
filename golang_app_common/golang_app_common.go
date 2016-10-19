package golang_app_common

import (
	"log"
)

func CheckError(worker_name string, message string, err error) {
	if err != nil {
		log.Fatalf("%s: %s: %s", worker_name, message, err)
	}
}

