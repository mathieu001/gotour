package cmd

import "github.com/spf13/cobra"

var rootCmd=&cobra.Command{
	Use:"",
	Short:"",
	Long:"",
}

func Execue() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(timeCmd)
}