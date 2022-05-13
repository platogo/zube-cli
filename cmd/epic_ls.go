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

	"github.com/InVisionApp/tabular"
	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
)

// epicLsCmd represents the epic ls command
var epicLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("epicLs called")
		if client, err := zube.NewClient(); err == nil {
			projectId, _ := cmd.Flags().GetInt("project-id")
			epics := client.FetchEpics(projectId)
			printEpics(&epics)
		}
	},
}

func init() {
	epicCmd.AddCommand(epicLsCmd)
	epicLsCmd.Flags().Int("project-id", 0, "Project ID")
	epicLsCmd.MarkFlagRequired("project-id")
}

func printEpics(epics *[]models.Epic) {
	tab := tabular.New()

	tab.Col("id", "ID", 6)
	tab.Col("title", "Title", 40)
	tab.Col("status", "Status", 10)

	format := tab.Print("id", "title", "status")
	for _, epic := range *epics {
		fmt.Printf(format, BrightMagenta(epic.Id), epic.Title, utils.SnakeCaseToTitleCase(epic.Status))
	}
}
