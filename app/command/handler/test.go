package handler

import (
	"fmt"
	"os"

	"lj/app/domain"
	"lj/app/execute"
	"lj/app/usecase"

	"github.com/spf13/cobra"
)

func Test(cmd *cobra.Command, args []string) {
	e := execute.NewExecute(os.Stdin, os.Stdout, os.Stderr)
	d := domain.NewDomain(e)
	u := usecase.NewUsecase(d)

	question := args[0]
	t, err := cmd.Flags().GetString("testcase")
	if err != nil {
		os.Exit(runtimeError)
	}

	if t != "" {
		if err := u.Judge(question, t); err != nil {
			fmt.Printf("ERROR: %s\n", err)
			os.Exit(failed)
		} else {
			os.Exit(success)
		}
	}

	if err := u.JudgeAll(question); err == usecase.ErrCanceled {
		os.Exit(runtimeError)
	} else if err != nil {
		os.Exit(failed)
	} else {
		fmt.Println("passed all test cases")
		os.Exit(success)
	}
}
