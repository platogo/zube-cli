/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/platogo/zube-cli/zube"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List cards with given filters",
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

		cards := client.FetchCards()

		for _, card := range cards {
			fmt.Printf("%s \n", card.Title)
		}
	},
}

func init() {
	cardCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
