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
var currentPersonCmd = &cobra.Command{
	Use:   "currentPerson",
	Short: "Show info about your own user",
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := zube.ParseDefaultConfig()
		if err != nil {
			log.Fatal(err)
			return
		}

		if isExp, _ := profile.IsAccessTokenExpired(); isExp {
			log.Fatal("Access Token is expired!")
			// TODO: If the token is expired, attempt to refresh it and resave it into the profile
			return
		}

		// Construct client
		client := zube.NewClientWithAccessToken(profile.ClientId, profile.AccessToken)
		// Call public client API to fetch resource that is needed, then print formatted output
		fmt.Printf("%+v", client.FetchCurrentPerson())
	},
}

func init() {
	rootCmd.AddCommand(currentPersonCmd)
}
