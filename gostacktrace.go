package gostacktrace

import (
	"runtime"
)

type (
	programCounter = uintptr
	pCounters      = []programCounter
)

func Get(skip uint) StackTrace {
	stackTrace := StackTrace{}
	programCounters := make(pCounters, 32)
	length := runtime.Callers(int(skip), programCounters)
	if len(programCounters) == 0 {
		return stackTrace
	}

	stackTrace.pCounters = programCounters[:length]
	return stackTrace
}

func getFrame(pc programCounter) Frame {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return Frame{
			File: "unknown",
		}
	}
	file, line := fn.FileLine(pc)
	return Frame{
		Func: fn.Name(),
		File: file,
		Line: uint(line),
	}
}
