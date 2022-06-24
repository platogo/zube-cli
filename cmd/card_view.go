/*
Copyright © 2022 Daniils Petrovs <daniils@platogo.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
)

// cardViewCmd represents the view command
var cardViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display the title, status, body and other info about a Zube card.",
	Run: func(cmd *cobra.Command, args []string) {
		var cardNumber string

		if len(args) > 0 {
			cardNumber = args[0]
		}

		client, _ := zube.NewClient()

		if parentCmd := cmd.Parent().Name(); parentCmd == "card" {
			cardQueryByNumber := zube.Query{Filter: zube.Filter{Where: map[string]any{"number": cardNumber}}}
			cards := client.FetchCards(&cardQueryByNumber)
			if len(cards) == 1 {
				card := cards[0]
				comments := client.FetchCardComments(card.Id)
				printCard(&card)
				printComments(&comments)
			} else {
				fmt.Println("Card not found!")
			}
		}
	},
}

func init() {
	cardCmd.AddCommand(cardViewCmd)
}

func printCard(card *models.Card) {
	var labels []string
	var assigneeNames []string

	for _, label := range card.Labels {
		labels = append(labels, label.Name)
	}

	for _, assignee := range card.Assignees {
		assigneeNames = append(assigneeNames, assignee.Username)
	}

	priority := card.Priority.OrElse(0)

	titleFormat := Reverse(card.Title + " #" + fmt.Sprint(card.Number)).Bold()
	statusFormat := Underline(utils.SnakeCaseToTitleCase(card.Status))
	bodyFormat := Gray(22, card.Body)

	fmt.Println(titleFormat)
	fmt.Println(statusFormat)
	fmt.Println(Bold("Assignees:"), strings.Join(assigneeNames, " "))
	fmt.Println(Bold("Labels:"), strings.Join(labels, " "))

	if priority != 0 {
		fmt.Println(Bold("Priority:"), fmt.Sprintf("P%d", priority))
	}

	fmt.Println()
	fmt.Println(bodyFormat)
	fmt.Println()
	fmt.Println(Bold("View this card on Zube: " + "https://zube.io/platogo/platogo/c/" + fmt.Sprint(card.Number))) // TODO: Replace with generic method
}

func printComments(comments *[]models.Comment) {

	fmt.Printf("------\n\n%s\n\n", Bold("Comments"))

	for _, comment := range *comments {
		fmt.Printf("%s\n%s\n\n", Reverse(comment.Creator.Name), Gray(14, comment.Timestamps.CreatedAt))

		fmt.Println(comment.Body)
	}
}