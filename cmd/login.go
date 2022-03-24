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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/platogo/zube-cli/zube"
	"github.com/spf13/cobra"
)

var ClientId string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		homedir, err := os.UserHomeDir()
		privateKeyFilePath, err := filepath.Abs(homedir + "/.ssh/zube_api_key.pem")
		privateKeyFile, err := ioutil.ReadFile(privateKeyFilePath)

		client := zube.NewClient(zube.ZubeHost, ClientId)

		if err != nil {
			log.Fatal(err)
		}

		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
