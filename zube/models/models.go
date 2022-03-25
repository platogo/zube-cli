package models

// Zube Response resources as defined in the official documentation: https://zube.io/docs/api

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	Total      int `json:"total"`
}

type Cards struct {
	Pagination
	Data []Card `json:"data"`
}

type ZubeAccessToken struct {
	AccessToken string `json:"access_token"`
}

type Person struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	AvatarPath string `json:"avatar_path"`
}

type CurrentPerson struct {
	Person
	GithubUserId int `json:"github_user_id"`
}

type Assignee struct {
	Person
}

type Epic struct {
	Id          int    `json:"id"`
	WorkspaceId int    `json:"workspace_id"`
	Number      int    `json:"number"`
	Status      string `json:"status"`
	Color       string `json:"color"`
	Title       string `json:"title"`
}

type Label struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	ProjectId int    `json:"project_id"`
}

type Milestone struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	SourceId int    `json:"source_id"`
	State    string `json:"state"`
}

type Source struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GithubIssue struct {
	Id          int       `json:"id"`
	MilestoneId int       `json:"milestone_id"`
	SourceId    int       `json:"source_id"`
	HtmlUrl     string    `json:"html_url"`
	Locked      bool      `json:"locked"`
	Merged      bool      `json:"merged"`
	Number      int       `json:"number"`
	Type        string    `json:"type"`
	CreatedAt   string    `json:"created_at"`
	Milestone   Milestone `json:"milestone"`
	Source      Source    `json:"source"`
}

type Card struct {
	Id            int        `json:"id"`
	CreatorId     int        `json:"creator_id"`
	ProjectId     int        `json:"project_id"`
	SprintId      int        `json:"sprint_id"`
	WorkspaceId   int        `json:"workspace_id"`
	Body          string     `json:"body"`
	CategoryName  string     `json:"category_name"`
	ClosedAt      string     `json:"closed_at"`
	CommentsCount int        `json:"comments_count"`
	LastCommentAt string     `json:"last_comment_at"`
	Number        int        `json:"number"`
	Points        int        `json:"points"`
	Priority      int        `json:"priority"`
	SearchKey     string     `json:"search_key"`
	State         string     `json:"state"`
	Status        string     `json:"status"`
	Title         string     `json:"title"`
	UpvotesCount  int        `json:"upvotes_count"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	EpicId        int        `json:"epic_id"`
	CloserId      int        `json:"closer_id"`
	Assignees     []Assignee `json:"assignees"`
	Creator       []Person   `json:"creator"`
	Epic          Epic       `json:"epic"`
	Labels        []Label    `json:"labels"`
}
