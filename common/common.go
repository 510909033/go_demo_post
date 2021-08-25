package common

import "log"

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrRecover() {
	if err := recover(); err != nil {
		log.Printf("recover:%+v", err)
	}
}
