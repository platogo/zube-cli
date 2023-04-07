package zube

import (
	"fmt"
	"log"
	"strconv"

	"github.com/markphelps/optional"
	"github.com/platogo/zube/models"
)

func Check(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}

// Convert projects to a slice of project names
func ProjectNames(projects *[]models.Project) []string {
	var names []string

	for _, p := range *projects {
		names = append(names, p.Name)
	}
	return names
}

func GetProjectByName(name string, projects *[]models.Project) (models.Project, error) {
	for _, p := range *projects {
		if p.Name == name {
			return p, nil
		}
	}
	return models.Project{}, fmt.Errorf("could not find project with name: %s", name)
}

// Convert workspaces to a slice of workspace names
func WorkspaceNames(workspaces *[]models.Workspace) []string {
	var names []string

	for _, w := range *workspaces {
		names = append(names, w.Name)
	}
	return names
}

func GetWorkspaceByName(name string, workspaces *[]models.Workspace) models.Workspace {
	for _, w := range *workspaces {
		if w.Name == name {
			return w
		}
	}
	return models.Workspace{}
}

// Convert labels to a slice of label names
func LabelNames(labels *[]models.Label) []string {
	var names []string

	for _, l := range *labels {
		names = append(names, l.Name)
	}
	return names
}

func LabelIds(labels *[]models.Label) []int {
	var ids []int

	for _, l := range *labels {
		ids = append(ids, l.Id)
	}
	return ids
}

func GetLabelsByIndexes(indexes []int, labels []models.Label) []models.Label {
	var res []models.Label

	for _, index := range indexes {
		res = append(res, labels[index])
	}
	return res
}

func EpicTitles(epics *[]models.Epic) []string {
	var titles []string

	for _, e := range *epics {
		titles = append(titles, e.Title)
	}

	return titles
}

func GetEpicByTitle(title string, epics *[]models.Epic) models.Epic {
	for _, e := range *epics {
		if e.Title == title {
			return e
		}
	}

	return models.Epic{}
}

func ParsePriority(priority string) optional.Int {
	if priority == "None" {
		return optional.Int{}
	}

	p, _ := strconv.Atoi(priority)

	return optional.NewInt(p)
}

func MemberNames(members *[]models.Member) []string {
	var names []string

	for _, m := range *members {
		names = append(names, m.Name)
	}

	return names
}

func GetMembersByNames(names []string, members []models.Member) []models.Member {
	var foundMembers []models.Member

	for _, m := range members {
		if Contains(names, m.Name) {
			foundMembers = append(foundMembers, m)
		}
	}

	return foundMembers
}

func MemberIds(members *[]models.Member) []int {
	var ids []int

	for _, m := range *members {
		ids = append(ids, m.Id)
	}

	return ids
}

func SourceNames(sources *[]models.Source) []string {
	var names []string

	for _, source := range *sources {
		names = append(names, source.Name)
	}

	return names
}

func GetSourceByName(name string, sources *[]models.Source) models.Source {
	for _, source := range *sources {
		if source.Name == name {
			return source
		}
	}

	return models.Source{}
}

// Check if `coll` contains an element `e` of type `T`
func Contains[T comparable](coll []T, e T) bool {
	for _, item := range coll {
		if item == e {
			return true
		}
	}
	return false
}

// Returns a direct URL to a Zube card
func CardUrl(account *models.Account, project *models.Project, card *models.Card) string {
	return fmt.Sprintf("https://zube.io/%s/%s/c/%d", account.Slug, project.Slug, card.Number)
}
