package zube

import (
	"errors"
	"fmt"

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
	return models.Project{}, errors.New(fmt.Sprintf("could not find project with name: %s", name))
}

// Convert workspaces to a slice of workspace names
func WorkspaceNames(workspaces *[]models.Workspace) []string {
	var names []string

	for _, w := range *workspaces {
		names = append(names, w.Name)
	}
	return names
}

func GetWorkspaceByName(name string, workspaces *[]models.Workspace) (models.Workspace, error) {
	for _, w := range *workspaces {
		if w.Name == name {
			return w, nil
		}
	}
	return models.Workspace{}, errors.New(fmt.Sprintf("could not find workspace with name: %s", name))
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
