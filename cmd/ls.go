/*
Copyright © 2022 Daniils Petrovs <daniils@platogo.com>

*/
package cmd

import (
	"fmt"
	"log"

	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List cards with given filters",
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := zube.ParseDefaultConfig()
		if err != nil {
			log.Fatal(err)
			return
		}

		client, _ := zube.NewClientWithProfile(&profile)

		if parentCmd := cmd.Parent().Name(); parentCmd == "card" {
			cards := client.FetchCards()
			printCards(&cards)
		}
	},
}

func init() {
	cardCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printCards(cards *[]models.Card) {
	const maxTitleWidth = 34
	formatString := "%-6d %" + fmt.Sprint(maxTitleWidth) + "s... %10s \n"
	// Print header
	fmt.Printf("%-6s %"+fmt.Sprint(maxTitleWidth+3)+"s %10s\n", Reverse(" No."), Reverse("Title              "), Reverse(" State  "))

	// Print rows
	for _, card := range *cards {
		fmt.Printf(formatString, BrightGreen(card.Number), BrightWhite(utils.TruncateString(card.Title, maxTitleWidth)), White(card.Status))
	}
}
