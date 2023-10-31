package utils

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/InVisionApp/tabular"
	. "github.com/logrusorgru/aurora"
	"github.com/platogo/zube"
	"github.com/platogo/zube/models"
	"github.com/samber/lo"
)

func PrintCards(cards *[]models.Card) {
	const maxTitleLen = 60

	tab := tabular.New()

	tab.Col("no", "Number", 6)
	tab.Col("title", "Title", maxTitleLen+6)
	tab.Col("status", "Status", 10)

	format := tab.Print("no", "title", "status")

	for _, card := range *cards {

		fmtTitle := TruncateString(card.Title, maxTitleLen)

		if utf8.RuneCountInString(card.Title) > maxTitleLen {
			fmtTitle += "..."
		}

		fmt.Printf(format,
			BrightGreen(card.Number),
			fmtTitle,
			SnakeCaseToTitleCase(card.Status),
		)
	}
}

func PrintCard(account *models.Account, project *models.Project, card *models.Card) {
	var labels []string
	var assigneeNames []string

	for _, label := range card.Labels {
		labels = append(labels, label.Name)
	}

	for _, assignee := range card.Assignees {
		assigneeNames = append(assigneeNames, assignee.Username)
	}

	priority := card.Priority.OrElse(0)

	titleFormat := Reverse(card.Title + " #" + fmt.Sprint(card.Number)).Bold()
	statusFormat := Underline(SnakeCaseToTitleCase(card.Status))
	bodyFormat := Gray(22, card.Body)
	cardUrl := zube.CardUrl(account, project, card)

	fmt.Println(titleFormat)
	fmt.Println(statusFormat)
	fmt.Println(Bold("Assignees:"), strings.Join(assigneeNames, " "))
	fmt.Println(Bold("Labels:"), strings.Join(labels, " "))

	if priority != 0 {
		fmt.Println(Bold("Priority:"), fmt.Sprintf("P%d", priority))
	}

	if card.GithubIssue.Id != 0 {
		fmt.Println(Bold("Github:"), fmt.Sprintf("%s#%d", card.GithubIssue.Source.Name, card.GithubIssue.Number))
	}

	fmt.Println()
	fmt.Println(bodyFormat)
	fmt.Println()
	fmt.Println(Bold("View this card on Zube: " + cardUrl))
}

func PrintComments(comments *[]models.Comment) {

	fmt.Printf("------\n\n%s\n\n", Bold("Comments"))

	for _, comment := range *comments {
		fmt.Printf("%s\n%s\n\n", Reverse(comment.Creator.Name), Gray(14, comment.Timestamps.CreatedAt))

		fmt.Println(comment.Body)
	}
}

func PrintEpics(epics *[]models.Epic) {
	tab := tabular.New()

	tab.Col("id", "ID", 6)
	tab.Col("title", "Title", 40)
	tab.Col("status", "Status", 10)

	format := tab.Print("id", "title", "status")
	for _, epic := range *epics {
		fmt.Printf(format, BrightMagenta(epic.Id), epic.Title, SnakeCaseToTitleCase(epic.Status))
	}
}

func PrintProjects(projects *[]models.Project) {
	tab := tabular.New()

	tab.Col("id", "ID", 4)
	tab.Col("name", "Name", 10)
	tab.Col("description", "Description", 20)

	format := tab.Print("id", "name", "description")
	for _, project := range *projects {
		fmt.Printf(format, BrightMagenta(project.Id), project.Name, project.Description)
	}
}

func PrintWorkspaces(workspaces *[]models.Workspace) {
	tab := tabular.New()

	tab.Col("id", "ID", 6)
	tab.Col("name", "Name", 20)
	tab.Col("description", "Description", 30)

	format := tab.Print("id", "name", "description")
	for _, workspace := range *workspaces {
		fmt.Printf(format,
			BrightYellow(workspace.Id),
			workspace.Name,
			workspace.Description,
		)
	}
}

func PrintSprints(sprints *[]models.Sprint) {
	tab := tabular.New()

	tab.Col("id", "ID", 6)
	tab.Col("title", "Title", 20)
	tab.Col("state", "State", 10)

	format := tab.Print("id", "title", "state")
	lo.ForEach(*sprints, func(sprint models.Sprint, _ int) {
		fmt.Printf(format, BrightYellow(sprint.Id), sprint.Title, sprint.State)
	})
}

func PrintSources(sources *[]models.Source) {
	tab := tabular.New()

	tab.Col("id", "ID", 6)
	tab.Col("name", "Name", 30)

	format := tab.Print("id", "name")
	lo.ForEach(*sources, func(source models.Source, _ int) {
		fmt.Printf(format, BrightYellow(source.Id), source.Name)
	})
}
