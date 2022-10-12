package models

import "github.com/markphelps/optional"

// Zube Response resources as defined in the official documentation: https://zube.io/docs/api

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	Total      int `json:"total"`
}

type Resource interface {
	Account | Card | Comment | Project | Sprint | Label | Epic | Source | Workspace | Member
}

type Data[T Resource] struct {
	Data []T
}

type PaginatedResponse[T Resource] struct {
	Pagination
	Data []T
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

type Member struct {
	Person
	GithubUserId       int  `json:"github_user_id"`
	IsUser             bool `json:"is_user"`
	NameIsLocked       bool `json:"name_is_locked"`
	AvatarPathIsLocked bool `json:"avatar_path_is_locked"`
	Timestamps
}

type Timestamps struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Account struct {
	Id          int    `json:"id"`
	DisplayName string `json:"display_name"`
	Status      string `json:"status"`
	Timestamps
	PrivateUsersCount int     `json:"private_users_count"`
	FirstBillableAt   string  `json:"first_billable_at"`
	Slug              string  `json:"slug"`
	HasGithubBilling  string  `json:"has_github_billing"`
	Discount          int     `json:"discount"`
	HasAnnualBilling  bool    `json:"has_annual_billing"`
	Seats             int     `json:"seats"`
	AnnualAmount      float64 `json:"annual_amount"`
}

type Epic struct {
	Id               int    `json:"id"`
	AssigneeId       int    `json:"assignee_id"`
	CreatorId        int    `json:"creator_id"`
	WorkspaceId      int    `json:"workspace_id"`
	ProjectId        int    `json:"project_id"`
	CardsStatus      string `json:"cards_status"`
	Description      string `json:"description"`
	DueOn            string `json:"due_on"`
	Number           int    `json:"number"`
	SearchKey        string `json:"search_key"`
	State            string `json:"state"`
	Status           string `json:"status"`
	Color            string `json:"color"`
	Title            string `json:"title"`
	TrackCards       bool   `json:"track_cards"`
	OpenCardsCount   int    `json:"open_cards_count"`
	ClosedCardsCount int    `json:"closed_cards_count"`
	OpenPoints       int    `json:"open_points"`
	ClosedPoints     int    `json:"closed_points"`
	CommentsCount    int    `json:"comments_count"`
	ClosedAt         string `json:"closed_at"`
	Timestamps
	CloserId   int    `json:"closed_id"`
	Priority   int    `json:"priority"`
	StartDate  string `json:"start_date"`
	Rank       string `json:"rank"`
	EpicListId int    `json:"epic_list_id"`
	Assignee   `json:"assignee"`
	Closer     Person  `json:"closer"`
	Creator    Person  `json:"creator"`
	Labels     []Label `json:"labels"`
}

type Workspace struct {
	Id             int    `json:"id"`
	ProjectId      int    `json:"project_id"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Private        bool   `json:"private"`
	PriorityFormat int    `json:"priority_format"`
	Priority       bool   `json:"priority"`
	Points         bool   `json:"points"`
	Upvotes        bool   `json:"upvotes"`
	Timestamps
	ArchiveMergedPrs        bool   `json:"archive_merged_prs"`
	UseCategoryLabels       bool   `json:"use_category_labels"`
	AutoArchiveClosedCards  bool   `json:"auto_archive_closed_cards"`
	ShouldUseFibonacciScale bool   `json:"should_use_fibonacci_scale"`
	Timezone                string `json:"timezone"`
	ShouldDisplayPrs        bool   `json:"should_display_prs"`
	DefaultCardTemplateId   int    `json:"default_card_template_id"`
}

type Project struct {
	Id          int    `json:"id"`
	AccountId   int    `json:"account_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Timestamps
	Slug                    string `json:"slug"`
	Private                 bool   `json:"private"`
	PriorityFormat          string `json:"priority_format"`
	Priority                bool   `json:"priority"`
	Points                  bool   `json:"points"`
	Triage                  bool   `json:"triage"`
	Upvotes                 bool   `json:"upvotes"`
	AutoAddGithubUsers      bool   `json:"auto_add_github_users"`
	Color                   string `json:"color"`
	ShouldUseFibonacciScale bool   `json:"should_use_fibonacci_scale"`
	DefaultEpicListId       int    `json:"default_epic_list_id"`
	Sources                 []Source
	Workspaces              []Workspace
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
	Id            int          `json:"id"`
	CreatorId     int          `json:"creator_id"`
	ProjectId     int          `json:"project_id"`
	SprintId      int          `json:"sprint_id"`
	WorkspaceId   int          `json:"workspace_id"`
	Body          string       `json:"body"`
	CategoryName  string       `json:"category_name"`
	ClosedAt      string       `json:"closed_at"`
	CommentsCount int          `json:"comments_count"`
	LastCommentAt string       `json:"last_comment_at"`
	Number        int          `json:"number"`
	Points        int          `json:"points"`
	Priority      optional.Int `json:"priority"` // must be one of 1, 2, 3, 4, 5, or null
	SearchKey     string       `json:"search_key"`
	State         string       `json:"state"`
	Status        string       `json:"status"`
	Title         string       `json:"title"`
	UpvotesCount  int          `json:"upvotes_count"`
	Timestamps
	GithubIssue GithubIssue `json:"github_issue"`
	EpicId      int         `json:"epic_id"`
	CloserId    int         `json:"closer_id"`
	Assignees   []Assignee  `json:"assignees"`
	AssigneeIds []int       `json:"assignee_ids"`
	Creator     []Person    `json:"creator"`
	Epic        Epic        `json:"epic"`
	Labels      []Label     `json:"labels"`
	LabelIds    []int       `json:"label_ids"`
}

type Comment struct {
	Id        int    `json:"id"`
	CardId    int    `json:"card_id"`
	CreatorId int    `json:"creator_id"`
	Body      string `json:"body"`
	Timestamps
	Creator Person
}

// Represents the parameters to be sent to Zube on card creation
// Only the `title` and `project_id` are required
type NewCardParameters struct {
	AssigneeIds  []int  `json:"assignee_ids"`
	Body         string `json:"body"`
	CategoryName string `json:"category_name"`
	EpicId       int    `json:"epic_id"`
	LabelIds     []int  `json:"label_ids"`
	Points       int    `json:"points"`
	Priority     int    `json:"priority"`   // must be one of 1, 2, 3, 4, 5, or null
	ProjectId    int    `json:"project_id"` // required
	SprintId     int    `json:"sprint_id"`
	Title        string `json:"title"` // required
	WorkspaceId  int    `json:"workspace_id"`
}

type Sprint struct {
	Id          int    `json:"id"`
	ClosedAt    string `json:"closed_at"`
	Description string `json:"description"`
	EndDate     string `json:"end_date"`
	StartDate   string `json:"start_date"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Timestamps
	ProjectId        int `json:"project_id"`
	WorkspaceId      int `json:"workspace_id"`
	OpenCardsCount   int `json:"open_cards_count"`
	ClosedCardsCount int `json:"closed_cards_count"`
	OpenPoints       int `json:"open_points"`
	ClosedPoints     int `json:"closed_points"`
}
