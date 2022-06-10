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
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the config init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the Zube CLI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		clientId := utils.StringPrompt("Enter your Zube Client ID:")

		if clientId == "" {
			fmt.Println(aurora.Red("Client ID cannot be blank!"))
			os.Exit(1)
		}

		viper.Set("client_id", clientId)
		viper.WriteConfig()
		fmt.Println(aurora.Green("Config initialized succesfully!"))
		fmt.Println("Don't forget to place your Zube private key at ~/.ssh/zube_api_key.pem")
	},
}

func init() {
	configCmd.AddCommand(initCmd)
}
