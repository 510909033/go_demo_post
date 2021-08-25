package my_groutinue_cpu_only_one

import (
	"go_demo_post/my_groutinue/lower_struct"
	//"github.com/felixge/fgprof"
	"log"
	"runtime"
)

func DemoMyGroutinueCpuOnlyOne() {

	lowerStruct := lower_struct.NewLowerStruct()
	lowerStruct.GetId()

	//var lower2 lower_struct.Lower2
	//log.Printf("low2.p=%p", &lower2)
	//lower2.Id = 2
	//lower2.GetId()
	//lower2.GetId2()

	var lower3 = &lower_struct.Lower2{Id: 22}
	log.Printf("lower3.p=%p", lower3)
	log.Printf("lower3。2.p=%p", &(*lower3))
	lower3.GetId()
	lower3.GetId2()
	return

	var s stackCounter
	s.Update()
	log.Println("over")
	return

	num := 1
	before := runtime.GOMAXPROCS(num)
	log.Printf("默认的runtime.GOMAXPROCS=%d", before)
	log.Printf("将runtime.GOMAXPROCS设置为 = %d", num)

	for i := 0; i < 10; i++ {
		go demoGrou()

	}

	runtime.Gosched()

	log.Printf("groutinue数量=%d", runtime.NumGoroutine())

	runtime.Breakpoint()

	log.Println("over")

	//pprof.StartCPUProfile(nil)
	//pprof.NewProfile().Add()
	//pprof.Profiles()

	//fgprof.Start()
}

func demoGrou() {
	log.Println("我是一个函数")
	runtime.Gosched()
}

type stackCounter map[string]int

func (s stackCounter) Update() {
	// Determine the runtime.Frame of this func so we can hide it from our
	// profiling output.
	rpc := make([]uintptr, 11)
	//rpc := make([]uintptr, 1)
	n1 := runtime.Callers(1, rpc)
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

	//fmt.Fprintf()
	return
	/*
	   	n := runtime.Callers(1, rpc)
	   	if n < 1 {
	   		panic("could not determine selfFrame")
	   	}
	   	selfFrame, _ := runtime.CallersFrames(rpc).Next()

	   	// COPYRIGHT: The code for populating `p` below is copied from
	   	// writeRuntimeProfile in src/runtime/pprof/pprof.go.
	   	//
	   	// Find out how many records there are (GoroutineProfile(nil)),
	   	// allocate that many records, and get the data.
	   	// There's a race—more records might be added between
	   	// the two calls—so allocate a few extra records for safety
	   	// and also try again if we're very unlucky.
	   	// The loop should only execute one iteration in the common case.
	   	var p []runtime.StackRecord
	   	n, ok := runtime.GoroutineProfile(nil)
	   	for {
	   		// Allocate room for a slightly bigger profile,
	   		// in case a few more entries have been added
	   		// since the call to ThreadProfile.
	   		p = make([]runtime.StackRecord, n+10)
	   		n, ok = runtime.GoroutineProfile(p)
	   		if ok {
	   			p = p[0:n]
	   			break
	   		}
	   		// Profile grew; try again.
	   	}

	   outer:
	   	for _, pp := range p {
	   		frames := runtime.CallersFrames(pp.Stack())

	   		var stack []string
	   		for {
	   			frame, more := frames.Next()
	   			if !more {
	   				break
	   			} else if frame.Entry == selfFrame.Entry {
	   				continue outer
	   			}

	   			stack = append([]string{frame.Function}, stack...)
	   		}
	   		key := strings.Join(stack, ";")
	   		s[key]++
	   	}

	*/
}
