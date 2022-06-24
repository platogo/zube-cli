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

	"github.com/AlecAivazis/survey/v2"
	"github.com/platogo/zube-cli/zube"
	"github.com/platogo/zube-cli/zube/models"
	"github.com/spf13/cobra"
)

// cardCreateCmd represents the create command
var cardCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Zube card",
	Long:  `Create a brand new Zube card for a given project.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, _ := zube.NewClient()

		projects := client.FetchProjects()
		workspaces := client.FetchWorkspaces()

		// We need to get the project ID before any other question, since the other prompt option fetchers
		// rely on it
		var projectName string
		projectPrompt := &survey.Select{
			Message: "Choose a project:",
			Options: zube.ProjectNames(&projects),
			Default: projects[0].Name,
		}

		survey.AskOne(projectPrompt, &projectName)
		project, err := zube.GetProjectByName(projectName, &projects)
		if err != nil {
			log.Fatalf(err.Error())
		}

		labels := client.FetchLabels(project.Id)
		epics := client.FetchEpics(project.Id)

		qs := []*survey.Question{
			{
				Name: "workspace",
				Prompt: &survey.Select{
					Message:  "Choose a workspace:",
					Options:  zube.WorkspaceNames(&workspaces),
					Default:  workspaces[0].Name,
					PageSize: 10,
				},
			},
			{
				Name:      "title",
				Prompt:    &survey.Input{Message: "Card title?"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name:   "description",
				Prompt: &survey.Multiline{Message: "Card description?"},
			},
			{
				Name: "labels",
				Prompt: &survey.MultiSelect{
					Message: "Choose labels:",
					Options: zube.LabelNames(&labels),
				},
			},
			{
				Name: "epic",
				Prompt: &survey.Select{
					Message: "Choose epic:",
					Options: append(zube.EpicTitles(&epics), "None"),
					Default: "None",
				},
			},
			{
				Name: "priority",
				Prompt: &survey.Select{
					Message: "Priority:",
					Options: []string{"None", "1", "2", "3", "4", "5"},
					Default: "None",
				},
			},
		}

		// TODO: Set assignees

		answers := struct {
			Workspace, Epic, Priority, Title, Description string
			Labels                                        []int
		}{}

		err = survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		workspace := zube.GetWorkspaceByName(answers.Workspace, &workspaces)

		epic := zube.GetEpicByTitle(answers.Epic, &epics)

		priority := zube.ParsePriority(answers.Priority)

		labels = zube.GetLabelsByIndexes(answers.Labels, labels)

		card := models.Card{
			ProjectId:   project.Id,
			WorkspaceId: workspace.Id,
			EpicId:      epic.Id,
			Title:       answers.Title,
			Priority:    priority,
			Body:        answers.Description,
			LabelIds:    zube.LabelIds(&labels)}

		respCard := client.CreateCard(&card)
		fmt.Printf("Card created with No: %d\n", respCard.Number)
	},
}

func init() {
	cardCmd.AddCommand(cardCreateCmd)
}
