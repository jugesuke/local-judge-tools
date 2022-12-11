package domain

import (
	"errors"
	"lj/app/execute"
	"lj/testcases"
)

type JudgeResult string

var (
	ErrCanceled = errors.New("canceled")
)

const (
	CE  JudgeResult = "CE"  // Compilation Error
	RE  JudgeResult = "RE"  // Runtime Error
	TLE JudgeResult = "TLE" // Time Limit Exceeded
	IE  JudgeResult = "IE"  // Internal Error
	WA  JudgeResult = "WA"  // Wrong Answer

	AC JudgeResult = "AC" // Accepted

	WJ JudgeResult = "WJ" // Waiting for Judge
)

func (j JudgeResult) Error() string {
	return string(j)
}

func (d *Domain) Judge(question string, testcase *testcases.Testcase) (JudgeResult, error) {
	out, err := d.execute.Test(question, testcase.GetStdin())
	if err == execute.ErrTimeout {
		return TLE, TLE
	} else if err == execute.ErrAborted {
		return RE, ErrCanceled
	} else if err != nil {
		return RE, err
	}
	if testcase.Is(out) {
		return AC, nil
	} else {
		return WA, WA
	}
}
