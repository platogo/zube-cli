/*
Copyright © 2022 Daniils Petrovs <daniils@platogo.com>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/InVisionApp/tabular"
	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// lsCmd represents the ls command
// Used to list various Zube entities, depending on the parent command name
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List cards with given filters",
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := zube.ParseDefaultConfig()
		if err != nil {
			log.Fatal(err)
			return
		}

		client, _ := zube.NewClientWithProfile(&profile)
		parentCmd := cmd.Parent().Name()

		switch parentCmd {
		case "card":
			query := newQueryFromFlags(cmd.LocalFlags())
			cards := client.FetchCards(&query)
			printCards(&cards)
		case "project":
			projects := client.FetchProjects()
			printProjects(&projects)
		}
	},
}

func init() {
	cardCmd.AddCommand(lsCmd)
	projectCmd.AddCommand(lsCmd)

	lsCmd.Flags().Int("id", 0, "Filter by card internal ID")
	lsCmd.Flags().String("category", "", "Filter by category name")
	lsCmd.Flags().Int("epic-id", 0, "Filter by epic ID")
	lsCmd.Flags().Int("number", 0, "Filter by card number")
	lsCmd.Flags().Int("priority", -1, "Filter by priority")
	lsCmd.Flags().Int("project-id", 0, "Filter by project ID")
	lsCmd.Flags().Int("sprint-id", 0, "Filter by sprint ID")
	lsCmd.Flags().String("state", "", "Filter by card state")
	lsCmd.Flags().String("status", "", "Filter by card status")
}

func printCards(cards *[]models.Card) {
	tab := tabular.New()

	tab.Col("no", "Number", 6)
	tab.Col("title", "Title", 46)
	tab.Col("status", "Status", 10)

	format := tab.Print("no", "title", "status")
	for _, card := range *cards {
		fmt.Printf(format,
			BrightGreen(card.Number),
			utils.TruncateString(card.Title, 40)+"...",
			utils.SnakeCaseToTitleCase(card.Status),
		)
	}
}

func printProjects(projects *[]models.Project) {
	tab := tabular.New()

	tab.Col("id", "ID", 4)
	tab.Col("name", "Name", 10)
	tab.Col("description", "Description", 20)

	format := tab.Print("id", "name", "description")
	for _, project := range *projects {
		fmt.Printf(format, BrightMagenta(project.Id), project.Name, project.Description)
	}
}

func newQueryFromFlags(flags *pflag.FlagSet) zube.Query {
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

	state, ok := flags.GetString("state")
	if ok == nil && state != "" {
		where["state"] = state
	}

	status, ok := flags.GetString("status")
	if ok == nil && status != "" {
		where["status"] = status
	}
	selectedCols := [4]string{"number", "title", "status", "category_name"}
	filter := zube.Filter{Where: where, Select: selectedCols[:]}

	return zube.Query{Filter: filter}
}
