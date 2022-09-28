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

		projects := client.FetchProjects(&zube.Query{})
		workspaces := client.FetchWorkspaces(&zube.Query{})
		sources := client.FetchSources()

		// We need to get the project ID before any other question, since the other prompt option fetchers
		// rely on it
		var projectName string
		projectPrompt := &survey.Select{
			Message: "Project:",
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
		members := client.FetchProjectMembers(project.Id)

		qs := []*survey.Question{
			{
				Name: "workspace",
				Prompt: &survey.Select{
					Message:  "Workspace:",
					Options:  zube.WorkspaceNames(&workspaces),
					Default:  workspaces[0].Name,
					PageSize: 10,
				},
			},
			{
				Name:      "title",
				Prompt:    &survey.Input{Message: "Title?"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name:   "description",
				Prompt: &survey.Editor{Message: "Description?", FileName: "*.md"},
			},
			{
				Name: "labels",
				Prompt: &survey.MultiSelect{
					Message: "Choose labels:",
					Options: zube.LabelNames(&labels),
				},
			},
			{
				Name: "assignees",
				Prompt: &survey.MultiSelect{
					Message: "Assignees:",
					Options: zube.MemberNames(&members),
				},
			},
			{
				Name: "epic",
				Prompt: &survey.Select{
					Message: "Epic:",
					Options: append(zube.EpicTitles(&epics), "None"),
					Default: "None",
				},
			},
			{
				Name: "source",
				Prompt: &survey.Select{
					Message: "Github source:",
					Options: append(zube.SourceNames(&sources), "None"),
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

		answers := struct {
			Workspace, Epic, Priority, Title, Description, Source string
			Labels                                                []int
			Assignees                                             []string
		}{}

		err = survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		workspace := zube.GetWorkspaceByName(answers.Workspace, &workspaces)

		epic := zube.GetEpicByTitle(answers.Epic, &epics)

		source := zube.GetSourceByName(answers.Source, &sources)

		priority := zube.ParsePriority(answers.Priority)

		labels = zube.GetLabelsByIndexes(answers.Labels, labels)

		assignees := zube.GetMembersByNames(answers.Assignees, members)

		card := models.Card{
			ProjectId:   project.Id,
			WorkspaceId: workspace.Id,
			EpicId:      epic.Id,
			Title:       answers.Title,
			Priority:    priority,
			Body:        answers.Description,
			LabelIds:    zube.LabelIds(&labels),
			AssigneeIds: zube.MemberIds(&assignees),
			GithubIssue: models.GithubIssue{SourceId: source.Id}}

		newCard := client.CreateCard(&card)
		account := client.FetchAccounts(
			&zube.Query{
				Filter: zube.Filter{Where: map[string]any{"id": project.AccountId}}})[0]

		fmt.Printf("\nView card on Zube: %s\n", zube.CardUrl(&account, &project, &newCard))
	},
}

func init() {
	cardCmd.AddCommand(cardCreateCmd)
}
