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

	"github.com/platogo/zube"
	"github.com/spf13/cobra"
)

var ClientId string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Zube with your client ID and private key.",
	Long:  `A command for debugging the login flow to Zube. On success, it will print your access token.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := zube.NewClientWithId(ClientId)

		privateKey, err := zube.GetPrivateKey()

		if err != nil {
			log.Fatal(err)
		}

		_, err = client.RefreshAccessToken(privateKey)

		fmt.Println("Access token:", client.AccessToken)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&ClientId, "client-id", "", "", "User's unique Zube Client ID")
	loginCmd.MarkPersistentFlagRequired("client-id")
}
