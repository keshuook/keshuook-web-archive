package githubapi

import (
	"context"
	"fmt"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/google/go-github/v80/github"
)

func formatCommits(raw []*github.RepositoryCommit) []Commits {
	var batch []Commits
	for _, x := range raw {
		batch = append(batch, Commits{
			SHA:  x.GetSHA(),
			Time: x.GetCommit().GetAuthor().GetDate().Time,
		})
	}
	return batch
}

func getCommitPage(client *github.Client, pageNumber int, repo *Repo) ([]*github.RepositoryCommit, *github.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	commits, resp, err := client.Repositories.ListCommits(ctx, repo.User, repo.RepoName, &github.CommitsListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    pageNumber,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return commits, resp, err
	}

	return commits, resp, nil
}

var CommitsList []Commits

func GetCommits(repo Repo) {
	client := github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_AUTH_TOKEN"))

	wg := sync.WaitGroup{}
	c := make(chan []Commits)
	errC := make(chan error)

	page1Commits, resp, err := getCommitPage(client, 1, &repo)

	if err != nil {
		fmt.Println("Critical Error: Failed to fetch first page!\n", err.Error())
	}

	allCommits := formatCommits(page1Commits)

	for i := 2; i <= resp.LastPage; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			rawCommits, _, err := getCommitPage(client, page, &repo)
			c <- formatCommits(rawCommits)
			if err != nil {
				errC <- err
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
		close(errC)
	}()

	for commit := range c {
		allCommits = append(allCommits, commit...)
	}
	for e := range errC {
		fmt.Println("Worker Error:", e)
	}

	slices.SortFunc(allCommits, func(a, b Commits) int {
		if a.Time.After(b.Time) {
			return -1
		}
		if a.Time.Before(b.Time) {
			return 1
		}
		return 0
	})

	CommitsList = allCommits
}
