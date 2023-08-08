package errlog

import "log"

func LogOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
