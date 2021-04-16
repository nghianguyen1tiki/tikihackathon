package assert

import "log"

func NotZero(value int, msg string) {
	if value == 0 {
		log.Panic(msg)
	}
}
