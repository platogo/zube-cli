/*
Copyright © 2022 Daniils Petrovs <daniils@platogo.com>

*/
package cmd

import (
	"fmt"
	"unicode/utf8"

	"github.com/InVisionApp/tabular"
	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Used to list various Zube entities, depending on the parent command name
var cardLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List cards with given filters",
	Run: func(cmd *cobra.Command, args []string) {
		profile := zube.Profile{ClientId: viper.GetString("client_id"), AccessToken: viper.GetString("access_token")}

		client, _ := zube.NewClientWithProfile(&profile)

		query := newQueryFromFlags(cmd.LocalFlags())

		var cards []models.Card

		if projectId, err := cmd.Flags().GetInt("project-id"); err == nil && projectId != 0 {
			query.Direction = "desc"
			query.Order.By = "milestone"
			cards = client.FetchProjectCards(projectId, &query)
		} else {
			cards = client.FetchCards(&query)
		}

		printCards(&cards)
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

func printCards(cards *[]models.Card) {
	tab := tabular.New()

	tab.Col("no", "Number", 6)
	tab.Col("title", "Title", 46)
	tab.Col("status", "Status", 10)

	format := tab.Print("no", "title", "status")

	const maxTitleLen = 40

	for _, card := range *cards {

		fmtTitle := utils.TruncateString(card.Title, maxTitleLen)

		if utf8.RuneCountInString(card.Title) > maxTitleLen {
			fmtTitle += "..."
		}

		fmt.Printf(format,
			BrightGreen(card.Number),
			fmtTitle,
			utils.SnakeCaseToTitleCase(card.Status),
		)
	}
}

func newQueryFromFlags(flags *pflag.FlagSet) zube.Query {
	var query zube.Query
	where := make(map[string]any)

	id, ok := flags.GetInt("id")
	if ok == nil && id != 0 {
		where["id"] = id
	}

	category, ok := flags.GetString("category")
	if ok == nil && category != "" {
		where["category_name"] = category
	}

	if epicId, ok := flags.GetInt("epic-id"); ok == nil && epicId != 0 {
		where["epic_id"] = epicId
	}

	if number, ok := flags.GetInt("number"); ok == nil && number != 0 {
		where["number"] = number
	}

	if priority, ok := flags.GetInt("priority"); ok == nil && priority >= 0 {
		where["priority"] = priority
	}

	if projectId, ok := flags.GetInt("project-id"); ok == nil && projectId != 0 {
		where["project_id"] = projectId
	}

	if sprintId, ok := flags.GetInt("sprint-id"); ok == nil && sprintId != 0 {
		where["sprint_id"] = sprintId
	}

	if workspaceId, ok := flags.GetInt("workspace-id"); ok == nil && workspaceId != 0 {
		where["workspace_id"] = workspaceId
	}

	if assigneeId, ok := flags.GetString("assignee-id"); ok == nil && assigneeId != "" {
		where["assignee_ids"] = []string{assigneeId}
	}

	state, ok := flags.GetString("state")
	if ok == nil && state != "" {
		where["state"] = state
	}

	status, ok := flags.GetString("status")
	if ok == nil && status != "" {
		where["status"] = status
	}
	selectedCols := [4]string{"number", "title", "status", "category_name"}
	query.Filter = zube.Filter{Where: where, Select: selectedCols[:]}

	return query
}
