package p

import (
    "fmt"
    "runtime"
)

func Print(f string, args ...interface{}) {
    pc, _, line, _ := runtime.Caller(1)
    fn := runtime.FuncForPC(pc)
    funcName := fn.Name()

    f = "%s:%d: " + f + "\n"
    args = append([]interface{}{funcName, line}, args...)

    fmt.Printf(f, args...)
}
