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
	"github.com/platogo/zube"
	"github.com/platogo/zube-cli/internal/utils"
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
		utils.PrintProjects(&projects)
	},
}

func init() {
	projectCmd.AddCommand(projectLsCmd)
}
