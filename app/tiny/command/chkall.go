package cmd

import (
	"lj/app/tiny/command/handler"

	"github.com/spf13/cobra"
)

var chkAllCmd = &cobra.Command{
	Use:   "chkall [program_name]",
	Short: "test your program with all test cases",
	Long:  "Test your program with all test cases.",
	Args:  cobra.MinimumNArgs(1),
	Run:   handler.ChkAll,
}

func init() {
	rootCmd.AddCommand(chkAllCmd)
}
