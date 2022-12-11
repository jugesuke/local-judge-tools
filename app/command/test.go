package cmd

import (
	"lj/app/command/handler"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test [question name]",
	Aliases: []string{"t"},
	Short:   "test your code with all test cases",
	Long:    `Test your code with all test cases.`,
	Args:    cobra.MinimumNArgs(1),
	Run:     handler.Test,
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("testcase", "t", "", "test with specific testcase")
}
