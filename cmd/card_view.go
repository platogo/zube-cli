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

	"github.com/spf13/cobra"

	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
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

				projectQueryById := zube.Query{Filter: zube.Filter{Where: map[string]any{"id": card.ProjectId}}}
				projects := client.FetchProjects(&projectQueryById)
				if len(projects) > 0 {
					project := projects[0]
					accountQueryById := zube.Query{Filter: zube.Filter{Where: map[string]any{"id": project.AccountId}}}
					accounts := client.FetchAccounts(&accountQueryById)

					utils.PrintCard(&accounts[0], &project, &card)
					utils.PrintComments(&comments)
				}

			} else {
				fmt.Println("Card not found!")
			}
		}
	},
}

func init() {
	cardCmd.AddCommand(cardViewCmd)
}
