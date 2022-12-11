package cmd

import (
	"lj/app/tiny/command/handler"

	"github.com/spf13/cobra"
)

var chkCmd = &cobra.Command{
	Use:   "chk [program_name] [testcase_number]",
	Short: "test your program with specific test cases",
	Long:  "Test your program with specific test cases.",
	Args:  cobra.MinimumNArgs(2),
	Run:   handler.Chk,
}

func init() {
	rootCmd.AddCommand(chkCmd)
}
