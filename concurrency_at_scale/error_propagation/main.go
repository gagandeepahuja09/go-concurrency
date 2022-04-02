package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}

// "lowlevel" module
type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{(wrapError(err, err.Error()))}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

// "intermediate" module
type IntermediateErr struct {
	error
}

func runJob(id string) error {
	const binPath = "/bad/job/binary"
	isExecuatable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return err
	} else if isExecuatable {
		return wrapError(nil, "job binary is not executable")
	}
}

func main() {

}
