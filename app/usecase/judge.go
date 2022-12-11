package usecase

import (
	"errors"
	"fmt"
	"time"

	"lj/app/domain"
	"lj/app/package/scol"
	"lj/testcases"
)

var (
	successCol = scol.Color{R: 55, G: 160, B: 55}
	errorCol   = scol.Color{R: 214, G: 153, B: 68}
	waitingCol = scol.Color{R: 119, G: 119, B: 119}
)

var (
	ErrFailed           = errors.New("failed")
	ErrCanceled         = errors.New("canceled")
	ErrTestcaseNotFound = errors.New("test case not found")
	ErrBuildFailed      = errors.New("build was failed")
)

func (u *Usecase) Judge(question string, testcaseName string) error {
	if err := u.domain.BuildWithoutWarning(question); err != nil {
		return ErrBuildFailed
	}
	t, err := testcases.GetTestcase(question, testcaseName)
	if err != nil {
		fmt.Printf("%s %s\n", question, scol.SetColor(scol.WHITE, errorCol, " "+string(domain.IE)+" "))
		fmt.Printf("ERROR: %s\n", ErrTestcaseNotFound)
		return ErrTestcaseNotFound
	}

	err = judge(u, question, t)
	if err != nil {
		return ErrFailed
	}
	return nil
}

func (u *Usecase) JudgeAll(question string) error {
	testcaseList, err := testcases.GetTestcases(question)
	if err != nil {
		fmt.Printf("%s %s\n", question, scol.SetColor(scol.WHITE, errorCol, " "+string(domain.IE)+" "))
		fmt.Printf("ERROR: %s\n", ErrTestcaseNotFound)
		return ErrTestcaseNotFound
	}

	if err := u.domain.BuildWithoutWarning(question); err != nil {
		fmt.Printf("%s %s\n", question, scol.SetColor(scol.WHITE, errorCol, " "+string(domain.CE)+" "))
		fmt.Printf("ERROR: %s\n", ErrBuildFailed)
		return ErrBuildFailed
	}

	var flag bool
	for _, t := range *testcaseList {
		err := judge(u, question, t)
		if err == domain.ErrCanceled {
			return ErrCanceled
		}
		if err != nil {
			flag = true
		}
	}
	if flag {
		return ErrFailed
	} else {
		return nil
	}
}

func judge(u *Usecase, question string, t *testcases.Testcase) error {
	cancel := make(chan struct{})
	go func(cancel chan struct{}) {
		var state int
		stateChar := [...]string{"|", "/", "-", "\\"}
		for {
			select {
			case <-cancel:
				return
			default:
				fmt.Printf("\x1b[2K\r%s %s %s", t.GetName(), scol.SetColor(scol.WHITE, waitingCol, " "+string(domain.WJ)+" "), stateChar[state])
				time.Sleep(100 * time.Millisecond)
				state = (state + 1) % len(stateChar)
			}
		}
	}(cancel)

	judgeResult, err := u.domain.Judge(question, t)
	close(cancel)

	if err == domain.ErrCanceled {
		fmt.Printf("\n")
		return domain.ErrCanceled
	} else if err != nil {
		fmt.Printf("\x1b[2K\r%s %s\n", t.GetName(), scol.SetColor(scol.WHITE, errorCol, " "+string(judgeResult)+" "))
		return judgeResult
	} else {
		fmt.Printf("\x1b[2K\r%s %s\n", t.GetName(), scol.SetColor(scol.WHITE, successCol, " AC "))
		return nil
	}
}
