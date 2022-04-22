/*
Copyright © 2022 Daniils Petrovs <daniils@platogo.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Manage cards",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("card called")
	},
}

func init() {
	rootCmd.AddCommand(cardCmd)
}
