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

	"github.com/InVisionApp/tabular"
	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
)

// projectLsCmd represents the projectLs command
var projectLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := zube.ParseDefaultConfig()
		if err != nil {
			log.Fatal(err)
			return
		}

		client, _ := zube.NewClientWithProfile(&profile)

		projects := client.FetchProjects()
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
