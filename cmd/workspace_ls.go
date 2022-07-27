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
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/spf13/cobra"
)

// workspaceLsCmd represents the workspace ls command
var workspaceLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		client, _ := zube.NewClient()
		workspaces := client.FetchWorkspaces(&zube.Query{})
		utils.PrintWorkspaces(&workspaces)
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceLsCmd)
}
