/*
Copyright © 2022 Daniils Petrovs <daniils@platogo.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const Version = "0.1.2"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "zube",
	Short:   "A Command Line utility for interacting with Zube.io",
	Long:    `Zube-CLI is a CLI tool built in Go that allows you to manage Zube cards, projects and other resources from the terminal.`,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zube-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
