package execute

import (
	"errors"
	"os"
)

type Execute struct {
	stdin  *os.File
	stdout *os.File
	stderr *os.File
}

var (
	ErrTimeout = errors.New("timeout")
	ErrAborted = errors.New("aborted")
)

func NewExecute(stdin, stdout, stderr *os.File) *Execute {
	return &Execute{stdin: stdin, stdout: stdout, stderr: stderr}
}
