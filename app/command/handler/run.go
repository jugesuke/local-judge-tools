package handler

import (
	"fmt"
	"lj/app/domain"
	"lj/app/execute"
	"lj/app/usecase"
	"os"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	e := execute.NewExecute(os.Stdin, os.Stdout, os.Stderr)
	d := domain.NewDomain(e)
	u := usecase.NewUsecase(d)

	question := args[0]

	if err := u.Run(question); err != nil {
		fmt.Println(err)
		os.Exit(failed)
	} else {
		os.Exit(success)
	}
}
