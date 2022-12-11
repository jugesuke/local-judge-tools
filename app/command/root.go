package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lj",
	Short: "local-judge-tools tests your code with test cases and helps you code",
	Long: `local-judge-tools is a CLI tool for checking your program with test cases.
	Also, local-judge-tools provide useful tools for coding.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
