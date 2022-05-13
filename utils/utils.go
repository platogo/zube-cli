package utils

import (
	"strings"

	"github.com/platogo/zube-cli/zube"
	"github.com/spf13/pflag"
)

// Returns a string truncated to the given `maxLen`
func TruncateString(s string, maxLen int) string {

	// Supports Japanese
	// Ref: Range loops https://blog.golang.org/strings
	// Source: https://dev.to/takakd/go-safe-truncate-string-9h0
	truncated := ""
	count := 0
	for _, char := range s {
		truncated += string(char)
		count++
		if count >= maxLen {
			break
		}
	}
	return truncated
}

// Converts a snake case string (e.g. `in_progress`) to title case (e.g. `In Progress`)
func SnakeCaseToTitleCase(s string) string {
	spacedWords := strings.ReplaceAll(s, "_", " ")
	return strings.Title(spacedWords)
}

// Constructs a zube `Query` from Cobra flags
func NewQueryFromFlags(flags *pflag.FlagSet) zube.Query {
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
