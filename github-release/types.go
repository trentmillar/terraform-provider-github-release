package github_release

type TotalCount struct {
	TotalCount int
}

type User struct {
	Id            string
	AvatarUrl     string
	Bio           string
	Company       string
	CreatedAt     string
	Email         string
	Followers     TotalCount
	Following     TotalCount
	Gists         TotalCount
	GistComments  TotalCount
	Issues        TotalCount
	IssueComments TotalCount
	Organizations TotalCount
	Repositories  TotalCount
}

type Release struct {
	DatabaseId    int
	Id            string
	Name          string
	Description   string
	TagName       string
	Url           string
	IsLatest      bool
	IsPrerelease  bool
	IsDraft       bool
	CreatedAt     string
	UpdatedAt     string
	PublishedAt   string
	Author        User
	Mentions      TotalCount
	ReleaseAssets TotalCount
}
