package usecase

import (
	"errors"
	"fmt"
	"lj/app/domain"
	"lj/testcases"
)

var (
	ErrFailed           = errors.New("failed")
	ErrCanceled         = errors.New("canceled")
	ErrTestcaseNotFound = errors.New("test case not found")
	ErrBuildFailed      = errors.New("build was failed")
)

func (u *TinyUsecase) Chk(question string, testcaseNumber int) error {
	testcaseList, err := testcases.GetTestcases(question)
	if err != nil {
		return ErrTestcaseNotFound
	}

	if len(*testcaseList) < testcaseNumber || testcaseNumber <= 0 {
		fmt.Printf("ERROR: %s\n", ErrTestcaseNotFound)
		return ErrTestcaseNotFound
	}

	if err := u.domain.BuildWithoutWarning(question); err != nil {
		fmt.Printf("ERROR: %s\n", ErrBuildFailed)
		return ErrBuildFailed
	}

	if err := judge(u, question, (*testcaseList)[testcaseNumber-1], testcaseNumber-1); err != nil {
		return ErrFailed
	} else {
		return nil
	}
}

func (u *TinyUsecase) ChkAll(question string) error {
	testcaseList, err := testcases.GetTestcases(question)
	if err != nil {
		fmt.Printf("ERROR: %s\n", ErrTestcaseNotFound)
		return ErrTestcaseNotFound
	}

	if err := u.domain.BuildWithoutWarning(question); err != nil {
		fmt.Printf("ERROR: %s\n", ErrBuildFailed)
		return ErrBuildFailed
	}

	var flag bool
	for i, t := range *testcaseList {
		err := judge(u, question, t, i)
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

func judge(u *TinyUsecase, question string, t *testcases.Testcase, testcasesNumber int) error {
	fmt.Printf("%02d ", testcasesNumber+1)
	judgeResult, err := u.domain.Judge(question, t)
	if err == domain.ErrCanceled {
		fmt.Printf("\n")
		return err
	} else if err != nil {
		fmt.Printf("%s\n", "NG")
		return judgeResult
	} else {
		fmt.Printf("%s\n", "OK")
		return nil
	}
}
