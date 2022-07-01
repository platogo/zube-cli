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
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
)

// projectLsCmd represents the projectLs command
var projectLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all Zube projects",
	Long:  `You can use this command to list all projects accessible to your user.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, _ := zube.NewClient()

		projects := client.FetchProjects(&zube.Query{})
		printProjects(&projects)
	},
}

func init() {
	projectCmd.AddCommand(projectLsCmd)
}

func printProjects(projects *[]models.Project) {
	tab := tabular.New()

	tab.Col("id", "ID", 4)
	tab.Col("name", "Name", 10)
	tab.Col("description", "Description", 20)

	format := tab.Print("id", "name", "description")
	for _, project := range *projects {
		fmt.Printf(format, BrightMagenta(project.Id), project.Name, project.Description)
	}
}
