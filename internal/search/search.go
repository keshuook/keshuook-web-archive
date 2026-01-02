package search

import (
	"errors"
	"time"

	"github.com/keshuook/keshuook-web-archive/internal/githubapi"
)

func Search(commits []githubapi.Commits, t time.Time) (string, error) {
	var low, high int = 0, len(commits)
	var mid int = (low + high) / 2

	for (high - low) > 1 {
		mid = (low + high) / 2
		if commits[mid].Time.After(t) {
			low = mid
		} else {
			high = mid
		}
	}

	if high >= len(commits) {
		return "", errors.New("No Commits Before the Given Time")
	}

	return commits[high].SHA, nil
}
