package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiny-lj",
	Short: "local-judge-tools (tiny version) tests your code with test cases",
	Long:  "local-judge-tools (tiny version) is a CLI tool for checking your program with test cases.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
