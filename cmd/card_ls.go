/*
Copyright © 2022 Daniils Petrovs <daniils@platogo.com>
*/
package cmd

import (
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
)

// Used to list various Zube entities, depending on the parent command name
var cardLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List cards with given filters",
	Run: func(cmd *cobra.Command, args []string) {
		client, _ := zube.NewClient()

		query := utils.NewQueryFromFlags(cmd.LocalFlags())

		var cards []models.Card

		if projectId, err := cmd.Flags().GetInt("project-id"); err == nil && projectId != 0 {
			query.Direction = "desc"
			query.Order.By = "milestone"
			cards = client.FetchProjectCards(projectId, &query)
		} else {
			cards = client.FetchCards(&query)
		}

		utils.PrintCards(&cards)
	},
}

func init() {
	cardCmd.AddCommand(cardLsCmd)

	cardLsCmd.Flags().Int("id", 0, "Filter by card internal ID")
	cardLsCmd.Flags().String("category", "", "Filter by category name")
	cardLsCmd.Flags().Int("epic-id", 0, "Filter by epic ID")
	cardLsCmd.Flags().Int("number", 0, "Filter by card number")
	cardLsCmd.Flags().Int("priority", -1, "Filter by priority")
	cardLsCmd.Flags().Int("project-id", 0, "Filter by project ID")
	cardLsCmd.Flags().Int("sprint-id", 0, "Filter by sprint ID")
	cardLsCmd.Flags().Int("workspace-id", 0, "Filter by workspace ID")
	cardLsCmd.Flags().String("assignee-id", "", "Filter by assignee")
	cardLsCmd.Flags().String("state", "", "Filter by card state")
	cardLsCmd.Flags().String("status", "", "Filter by card status")
}
