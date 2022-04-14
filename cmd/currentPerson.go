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

	"github.com/platogo/zube-cli/zube"
	"github.com/spf13/cobra"
)

// currentPersonCmd represents the currentPerson command
// This command is mostly for reference on how to build out the other commands
var currentPersonCmd = &cobra.Command{
	Use:   "currentPerson",
	Short: "Show info about your own user",
	Run: func(cmd *cobra.Command, args []string) {
		// Load any existing cached config
		profile, err := zube.ParseDefaultConfig()
		if err != nil {
			log.Fatalln(err)
			return
		}

		// Prepare a client using the loaded profile
		client, _ := zube.NewClientWithProfile(&profile)

		// Call public client API to fetch resource that is needed, then print formatted output
		person := client.FetchCurrentPerson()
		fmt.Printf("Username: %s\nName: %s\nId: %d\n", person.Username, person.Name, person.Id)
	},
}

func init() {
	rootCmd.AddCommand(currentPersonCmd)
}
