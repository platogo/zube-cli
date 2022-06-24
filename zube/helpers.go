package zube

import (
	"fmt"
	"strconv"

	"github.com/markphelps/optional"
	"github.com/platogo/zube-cli/zube/models"
)

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
