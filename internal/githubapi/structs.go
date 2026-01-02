package githubapi

import "time"

type Commits struct {
	SHA  string
	Time time.Time
}

type Repo struct {
	User     string
	RepoName string
}
