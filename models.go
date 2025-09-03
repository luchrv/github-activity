package main

import (
	"time"
)

// Event representa un evento de GitHub
type Event struct {
	ID        string    `json:"id"`
	Type      *string   `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Org       *Actor    `json:"org,omitempty"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

// Actor representa un actor de GitHub (usuario u organización)
type Actor struct {
	ID           int     `json:"id"`
	Login        string  `json:"login"`
	DisplayLogin string  `json:"display_login,omitempty"`
	GravatarID   *string `json:"gravatar_id"`
	URL          string  `json:"url"`
	AvatarURL    string  `json:"avatar_url"`
}

// Repo representa un repositorio de GitHub
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Payload representa el payload de un evento
type Payload struct {
	Action  *string         `json:"action,omitempty"`
	Issue   *Issue          `json:"issue,omitempty"`
	Comment *IssueComment   `json:"comment,omitempty"`
	Pages   []PageOperation `json:"pages,omitempty"`
}

// Issue representa un issue de GitHub
type Issue struct {
	ID                       int64                    `json:"id"`
	NodeID                   string                   `json:"node_id"`
	URL                      string                   `json:"url"`
	RepositoryURL            string                   `json:"repository_url"`
	LabelsURL                string                   `json:"labels_url"`
	CommentsURL              string                   `json:"comments_url"`
	EventsURL                string                   `json:"events_url"`
	HTMLURL                  string                   `json:"html_url"`
	Number                   int                      `json:"number"`
	State                    string                   `json:"state"`
	StateReason              *string                  `json:"state_reason,omitempty"`
	Title                    string                   `json:"title"`
	Body                     *string                  `json:"body,omitempty"`
	User                     *SimpleUser              `json:"user"`
	Labels                   []interface{}            `json:"labels"`
	Assignee                 *SimpleUser              `json:"assignee"`
	Assignees                []SimpleUser             `json:"assignees,omitempty"`
	Milestone                *Milestone               `json:"milestone"`
	Locked                   bool                     `json:"locked"`
	ActiveLockReason         *string                  `json:"active_lock_reason,omitempty"`
	Comments                 int                      `json:"comments"`
	PullRequest              *PullRequestReference    `json:"pull_request,omitempty"`
	ClosedAt                 *time.Time               `json:"closed_at,omitempty"`
	CreatedAt                time.Time                `json:"created_at"`
	UpdatedAt                time.Time                `json:"updated_at"`
	Draft                    *bool                    `json:"draft,omitempty"`
	ClosedBy                 *SimpleUser              `json:"closed_by,omitempty"`
	BodyHTML                 *string                  `json:"body_html,omitempty"`
	BodyText                 *string                  `json:"body_text,omitempty"`
	TimelineURL              *string                  `json:"timeline_url,omitempty"`
	Type                     *IssueType               `json:"type,omitempty"`
	Repository               *Repository              `json:"repository,omitempty"`
	PerformedViaGithubApp    interface{}              `json:"performed_via_github_app,omitempty"`
	AuthorAssociation        string                   `json:"author_association"`
	Reactions                ReactionRollup           `json:"reactions"`
	SubIssuesSummary         SubIssuesSummary         `json:"sub_issues_summary"`
	ParentIssueURL           *string                  `json:"parent_issue_url,omitempty"`
	IssueDependenciesSummary IssueDependenciesSummary `json:"issue_dependencies_summary"`
	IssueFieldValues         []IssueFieldValue        `json:"issue_field_values,omitempty"`
}

// SimpleUser representa un usuario simple de GitHub
type SimpleUser struct {
	Name              *string `json:"name,omitempty"`
	Email             *string `json:"email,omitempty"`
	Login             string  `json:"login"`
	ID                int64   `json:"id"`
	NodeID            string  `json:"node_id"`
	AvatarURL         string  `json:"avatar_url"`
	GravatarID        *string `json:"gravatar_id"`
	URL               string  `json:"url"`
	HTMLURL           string  `json:"html_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReposURL          string  `json:"repos_url"`
	EventsURL         string  `json:"events_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

// Milestone representa un hito de GitHub
type Milestone struct {
	URL          string      `json:"url"`
	HTMLURL      string      `json:"html_url"`
	LabelsURL    string      `json:"labels_url"`
	ID           int         `json:"id"`
	NodeID       string      `json:"node_id"`
	Number       int         `json:"number"`
	State        string      `json:"state"`
	Title        string      `json:"title"`
	Description  *string     `json:"description,omitempty"`
	Creator      *SimpleUser `json:"creator"`
	OpenIssues   int         `json:"open_issues"`
	ClosedIssues int         `json:"closed_issues"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	ClosedAt     *time.Time  `json:"closed_at,omitempty"`
	DueOn        *time.Time  `json:"due_on,omitempty"`
}

// PullRequestReference representa una referencia a un pull request
type PullRequestReference struct {
	MergedAt *time.Time `json:"merged_at,omitempty"`
	DiffURL  *string    `json:"diff_url,omitempty"`
	HTMLURL  *string    `json:"html_url,omitempty"`
	PatchURL *string    `json:"patch_url,omitempty"`
	URL      *string    `json:"url,omitempty"`
}

// IssueType representa el tipo de issue
type IssueType struct {
	ID          int       `json:"id"`
	NodeID      string    `json:"node_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Color       *string   `json:"color,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsEnabled   bool      `json:"is_enabled"`
}

// Repository representa un repositorio de GitHub
type Repository struct {
	ID                        int64                  `json:"id"`
	NodeID                    string                 `json:"node_id"`
	Name                      string                 `json:"name"`
	FullName                  string                 `json:"full_name"`
	License                   *LicenseSimple         `json:"license,omitempty"`
	Forks                     int                    `json:"forks"`
	Permissions               *Permissions           `json:"permissions,omitempty"`
	Owner                     SimpleUser             `json:"owner"`
	Private                   bool                   `json:"private"`
	HTMLURL                   string                 `json:"html_url"`
	Description               *string                `json:"description,omitempty"`
	Fork                      bool                   `json:"fork"`
	URL                       string                 `json:"url"`
	ArchiveURL                string                 `json:"archive_url"`
	AssigneesURL              string                 `json:"assignees_url"`
	BlobsURL                  string                 `json:"blobs_url"`
	BranchesURL               string                 `json:"branches_url"`
	CollaboratorsURL          string                 `json:"collaborators_url"`
	CommentsURL               string                 `json:"comments_url"`
	CommitsURL                string                 `json:"commits_url"`
	CompareURL                string                 `json:"compare_url"`
	ContentsURL               string                 `json:"contents_url"`
	ContributorsURL           string                 `json:"contributors_url"`
	DeploymentsURL            string                 `json:"deployments_url"`
	DownloadsURL              string                 `json:"downloads_url"`
	EventsURL                 string                 `json:"events_url"`
	ForksURL                  string                 `json:"forks_url"`
	GitCommitsURL             string                 `json:"git_commits_url"`
	GitRefsURL                string                 `json:"git_refs_url"`
	GitTagsURL                string                 `json:"git_tags_url"`
	GitURL                    string                 `json:"git_url"`
	IssueCommentURL           string                 `json:"issue_comment_url"`
	IssueEventsURL            string                 `json:"issue_events_url"`
	IssuesURL                 string                 `json:"issues_url"`
	KeysURL                   string                 `json:"keys_url"`
	LabelsURL                 string                 `json:"labels_url"`
	LanguagesURL              string                 `json:"languages_url"`
	MergetsURL                string                 `json:"merges_url"`
	MilestonesURL             string                 `json:"milestones_url"`
	NotificationsURL          string                 `json:"notifications_url"`
	PullsURL                  string                 `json:"pulls_url"`
	ReleasesURL               string                 `json:"releases_url"`
	SSHURL                    string                 `json:"ssh_url"`
	StargazersURL             string                 `json:"stargazers_url"`
	StatusesURL               string                 `json:"statuses_url"`
	SubscribersURL            string                 `json:"subscribers_url"`
	SubscriptionURL           string                 `json:"subscription_url"`
	TagsURL                   string                 `json:"tags_url"`
	TeamsURL                  string                 `json:"teams_url"`
	TreesURL                  string                 `json:"trees_url"`
	CloneURL                  string                 `json:"clone_url"`
	MirrorURL                 *string                `json:"mirror_url,omitempty"`
	HooksURL                  string                 `json:"hooks_url"`
	SvnURL                    string                 `json:"svn_url"`
	Homepage                  *string                `json:"homepage,omitempty"`
	Language                  *string                `json:"language,omitempty"`
	ForksCount                int                    `json:"forks_count"`
	StargazersCount           int                    `json:"stargazers_count"`
	WatchersCount             int                    `json:"watchers_count"`
	Size                      int                    `json:"size"`
	DefaultBranch             string                 `json:"default_branch"`
	OpenIssuesCount           int                    `json:"open_issues_count"`
	IsTemplate                bool                   `json:"is_template"`
	Topics                    []string               `json:"topics,omitempty"`
	HasIssues                 bool                   `json:"has_issues"`
	HasProjects               bool                   `json:"has_projects"`
	HasWiki                   bool                   `json:"has_wiki"`
	HasPages                  bool                   `json:"has_pages"`
	HasDownloads              bool                   `json:"has_downloads"`
	HasDiscussions            bool                   `json:"has_discussions"`
	Archived                  bool                   `json:"archived"`
	Disabled                  bool                   `json:"disabled"`
	Visibility                string                 `json:"visibility"`
	PushedAt                  *time.Time             `json:"pushed_at,omitempty"`
	CreatedAt                 *time.Time             `json:"created_at,omitempty"`
	UpdatedAt                 *time.Time             `json:"updated_at,omitempty"`
	AllowRebaseMerge          bool                   `json:"allow_rebase_merge"`
	TempCloneToken            *string                `json:"temp_clone_token,omitempty"`
	AllowSquashMerge          bool                   `json:"allow_squash_merge"`
	AllowAutoMerge            bool                   `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool                   `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool                   `json:"allow_update_branch"`
	UseSquashPRTitleAsDefault bool                   `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitTitle    *string                `json:"squash_merge_commit_title,omitempty"`
	SquashMergeCommitMessage  *string                `json:"squash_merge_commit_message,omitempty"`
	MergeCommitTitle          *string                `json:"merge_commit_title,omitempty"`
	MergeCommitMessage        *string                `json:"merge_commit_message,omitempty"`
	AllowMergeCommit          bool                   `json:"allow_merge_commit"`
	AllowForking              bool                   `json:"allow_forking"`
	WebCommitSignoffRequired  bool                   `json:"web_commit_signoff_required"`
	OpenIssues                int                    `json:"open_issues"`
	Watchers                  int                    `json:"watchers"`
	MasterBranch              *string                `json:"master_branch,omitempty"`
	StarredAt                 *string                `json:"starred_at,omitempty"`
	AnonymousAccessEnabled    *bool                  `json:"anonymous_access_enabled,omitempty"`
	CodeSearchIndexStatus     *CodeSearchIndexStatus `json:"code_search_index_status,omitempty"`
}

// LicenseSimple representa una licencia simple
type LicenseSimple struct {
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	URL     *string `json:"url,omitempty"`
	SpdxID  *string `json:"spdx_id,omitempty"`
	NodeID  string  `json:"node_id"`
	HTMLURL *string `json:"html_url,omitempty"`
}

// Permissions representa los permisos de un repositorio
type Permissions struct {
	Admin    bool `json:"admin"`
	Pull     bool `json:"pull"`
	Triage   bool `json:"triage,omitempty"`
	Push     bool `json:"push"`
	Maintain bool `json:"maintain,omitempty"`
}

// CodeSearchIndexStatus representa el estado del índice de búsqueda de código
type CodeSearchIndexStatus struct {
	LexicalSearchOK  bool   `json:"lexical_search_ok"`
	LexicalCommitSHA string `json:"lexical_commit_sha"`
}

// ReactionRollup representa un resumen de reacciones
type ReactionRollup struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	PlusOne    int    `json:"+1"`
	MinusOne   int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Hooray     int    `json:"hooray"`
	Eyes       int    `json:"eyes"`
	Rocket     int    `json:"rocket"`
}

// SubIssuesSummary representa un resumen de sub-issues
type SubIssuesSummary struct {
	Total            int `json:"total"`
	Completed        int `json:"completed"`
	PercentCompleted int `json:"percent_completed"`
}

// IssueDependenciesSummary representa un resumen de dependencias de issues
type IssueDependenciesSummary struct {
	BlockedBy      int `json:"blocked_by"`
	Blocking       int `json:"blocking"`
	TotalBlockedBy int `json:"total_blocked_by"`
	TotalBlocking  int `json:"total_blocking"`
}

// IssueFieldValue representa un valor de campo de issue
type IssueFieldValue struct {
	IssueFieldID       int64         `json:"issue_field_id"`
	NodeID             string        `json:"node_id"`
	DataType           string        `json:"data_type"`
	Value              interface{}   `json:"value"`
	SingleSelectOption *SelectOption `json:"single_select_option,omitempty"`
}

// SelectOption representa una opción de selección única
type SelectOption struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// IssueComment representa un comentario de issue
type IssueComment struct {
	ID                    int64          `json:"id"`
	NodeID                string         `json:"node_id"`
	URL                   string         `json:"url"`
	Body                  string         `json:"body"`
	BodyText              *string        `json:"body_text,omitempty"`
	BodyHTML              *string        `json:"body_html,omitempty"`
	HTMLURL               string         `json:"html_url"`
	User                  *SimpleUser    `json:"user"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	IssueURL              string         `json:"issue_url"`
	AuthorAssociation     string         `json:"author_association"`
	PerformedViaGithubApp interface{}    `json:"performed_via_github_app,omitempty"`
	Reactions             ReactionRollup `json:"reactions"`
}

// PageOperation representa una operación de página
type PageOperation struct {
	PageName string  `json:"page_name"`
	Title    string  `json:"title"`
	Summary  *string `json:"summary,omitempty"`
	Action   string  `json:"action"`
	SHA      string  `json:"sha"`
	HTMLURL  string  `json:"html_url"`
}

// Events representa un array de eventos
type Events []Event
