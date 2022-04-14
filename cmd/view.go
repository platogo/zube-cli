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
				printCard(&cards[0])
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

	fmt.Printf("%s\n%s\n%s\n\n%s",
		Reverse(card.Title+" #"+fmt.Sprint(card.Number)).Bold(),
		Underline(utils.SnakeCaseToTitleCase(card.Status)),
		Bold("Labels: "+strings.Join(labels, " ")),
		Gray(12, card.Body))
}
