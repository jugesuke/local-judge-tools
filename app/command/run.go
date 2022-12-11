/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lj/app/command/handler"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [question name]",
	Aliases: []string{"r"},
	Short:   "Build and Run your program",
	Long:    "Build your program, and Execute it",
	Args:    cobra.MinimumNArgs(1),
	Run:     handler.Run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}
