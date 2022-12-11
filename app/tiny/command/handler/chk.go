package handler

import (
	"fmt"
	"os"
	"strconv"

	"lj/app/domain"
	"lj/app/execute"
	"lj/app/tiny/usecase"

	"github.com/spf13/cobra"
)

func Chk(cmd *cobra.Command, args []string) {
	e := execute.NewExecute(os.Stdin, os.Stdout, os.Stderr)
	d := domain.NewDomain(e)
	u := usecase.NewTinyUsecase(d)

	question := args[0]

	var testcaseNumber int
	if n, err := strconv.ParseInt(args[1], 10, 32); err != nil {
		fmt.Println("ERROR: test case not found")
		os.Exit(runtimeError)
	} else {
		testcaseNumber = int(n)
	}

	if err := u.Chk(question, testcaseNumber); err != nil {
		os.Exit(failed)
	} else {
		os.Exit(success)
	}
}

func ChkAll(cmd *cobra.Command, args []string) {
	e := execute.NewExecute(os.Stdin, os.Stdout, os.Stderr)
	d := domain.NewDomain(e)
	u := usecase.NewTinyUsecase(d)

	question := args[0]

	if err := u.ChkAll(question); err == usecase.ErrCanceled {
		os.Exit(runtimeError)
	} else if err != nil {
		os.Exit(failed)
	} else {
		fmt.Println("successful execution for all test cases")
		os.Exit(success)
	}
}
