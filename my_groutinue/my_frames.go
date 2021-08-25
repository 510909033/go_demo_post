package my_groutinue

import (
	"log"
	"runtime"
)

func MyFrame() {

	for {
		rpc := make([]uintptr, 100)
		n1 := runtime.Callers(0, rpc)
		log.Println("n1=", n1)
		frames := runtime.CallersFrames(rpc) //.Next()
		log.Println(frames)

		for {
			if frame, more := frames.Next(); more {
				//log.Printf("frame=%#v", frame)
				log.Printf("frame.Line = %d ,frame.Function=%s, frame.File = %s", frame.Line, frame.Function, frame.File)
				log.Printf("frame.Func.Name()=%s", frame.Func.Name())
				file, line := frame.Func.FileLine(frame.PC)
				log.Printf("frame.Func.FileLine(frame.PC), file=%s , line=%d", file, line)

				file, line = frame.Func.FileLine(frame.Entry)
				log.Printf("frame.Func.FileLine(frame.Entry), file=%s , line=%d", file, line)

				log.Println("")
			} else {
				break
			}
		}
	}
}
