/*
Copyright © 2023 Daniils Petrovs <daniils@platogo.com>

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
	"log"

	"github.com/platogo/zube"
	"github.com/platogo/zube-cli/internal/utils"
	"github.com/platogo/zube/models"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var labelLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all Zube labels",
	Long:  `List all registered labels in a project. Will print the color of the label in supported terminals.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, _ := zube.NewClient()

		var labels []models.Label

		if projectId, err := cmd.Flags().GetInt("project-id"); err == nil && projectId != 0 {
			labels = client.FetchLabels(projectId)
		} else {
			log.Fatal("Project ID is required")
		}

		utils.PrintItems(&labels)
	},
}

func init() {
	labelCmd.AddCommand(labelLsCmd)

	labelLsCmd.Flags().Int("project-id", 0, "Filter by project ID")
}
