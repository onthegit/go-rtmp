package rtmp

import (
	"runtime"
	"strconv"
)

type stackEntry struct {
	FilePath string
	Func     string
	Line     int
}

func stackEntriesToString(newLine string, s ...stackEntry) string {
	str := ""
	for _, u := range s {
		str += u.FilePath + ":" + strconv.Itoa(u.Line) + newLine + "(" + u.Func + ")" + newLine
	}
	return str
}

// stackTraceString returns the stack in string format
func stackTraceString(skip int, newLine string) string {
	return stackEntriesToString(newLine, stackTrace(skip+1)...)
}

// stackTrace returns the stack
func stackTrace(skip int) []stackEntry {

	stack := make([]stackEntry, 0, 15)
	for {
		p, pth, ln, g := runtime.Caller(skip)
		if !g {
			break
		}

		stackEntry := stackEntry{
			FilePath: pth,
			Func:     runtime.FuncForPC(p).Name(),
			Line:     ln,
		}
		stack = append(stack, stackEntry)
		skip++
	}
	return stack
}
