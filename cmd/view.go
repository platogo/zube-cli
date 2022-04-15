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
	"log"
	"strings"

	"github.com/spf13/cobra"

	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
)

// TODO: Rename to cardViewCmd
// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display the title, status, body and other info about a Zube card.",
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := zube.ParseDefaultConfig()
		if err != nil {
			log.Fatal(err)
			return
		}

		var cardNumber string

		if len(args) > 0 {
			cardNumber = args[0]
		}

		client, _ := zube.NewClientWithProfile(&profile)

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
	cardCmd.AddCommand(viewCmd)
}

func printCard(card *models.Card) {
	var labels []string

	for _, label := range card.Labels {
		labels = append(labels, label.Name)
	}

	format := "%s\n%s\n%s\nPriority: P%d\n\n%s"
	titleFormat := Reverse(card.Title + " #" + fmt.Sprint(card.Number)).Bold()
	statusFormat := Underline(utils.SnakeCaseToTitleCase(card.Status))
	bodyFormat := Gray(22, card.Body)

	fmt.Printf(format,
		titleFormat,
		statusFormat,
		Bold("Labels: "+strings.Join(labels, " ")),
		card.Priority,
		bodyFormat)
}

func printComments(comments *[]models.Comment) {

	fmt.Printf("\n\n%s\n\n", Bold("Comments"))

	for _, comment := range *comments {
		fmt.Printf("%s\n%s\n\n", Reverse(comment.Creator.Name), Gray(14, comment.Timestamps.CreatedAt))

		fmt.Println(comment.Body)
	}
}
