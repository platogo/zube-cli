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
	"time"

	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Zube cards",
	Long:  `Search all Zube cards using a fuzzy search query.`,
	Run: func(cmd *cobra.Command, args []string) {
		var searchQuery string

		switch {
		case len(args) == 0:
			fmt.Println("please provide a search query")
		case len(args) > 1:
			searchQuery = strings.Join(args, " ")
		default:
			searchQuery = args[0]
		}

		client, _ := zube.NewClient()
		cards := client.SearchCards(&zube.Query{Search: searchQuery})

		switch len(cards) {
		case 0:
			fmt.Println("no results")
		case 1:
			card := cards[0]
			projectQueryById := zube.Query{Filter: zube.Filter{Where: map[string]any{"id": card.ProjectId}}}
			project := client.FetchProjects(&projectQueryById)[0]
			accountQueryById := zube.Query{Filter: zube.Filter{Where: map[string]any{"id": project.AccountId}}}
			account := client.FetchAccounts(&accountQueryById)[0]
			utils.PrintCard(&account, &project, &card)
		default:
			utils.PrintCards(&cards)
		}
		time.Sleep(time.Second)
	},
}

func init() {
	cardCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
