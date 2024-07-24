package errorUtils

import (
	"fmt"
	"runtime"
)

var funcInfoFormat = "Stack Trace: {%s:%d} [%s]"

func getFuncInfo(pc uintptr, file string, line int) string {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return fmt.Sprintf(funcInfoFormat, file, line, "unknwon")
	}
	return fmt.Sprintf(funcInfoFormat, file, line, f.Name())
}

var wrapFormat = "%w\n%s" // "error \n {file:line} [func name] msg"

func wrap(err error, msg string) error {
	pc, file, line, ok := runtime.Caller(2)

	if !ok {
		return fmt.Errorf(wrapFormat, err, msg)
	}

	// {file:line} [funcName] msg
	stack := fmt.Sprintf("%s %s", getFuncInfo(pc, file, line), msg)
	return fmt.Errorf(wrapFormat, err, stack)
}

func WrapWithMessage(err error, msg string) error {
	return wrap(err, msg)
}

func Wrap(err error) error {
	return wrap(err, "")
}
