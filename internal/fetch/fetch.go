package fetch

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/keshuook/keshuook-web-archive/internal/githubapi"
)

func Get(SHA string, path string, REPO *githubapi.Repo) ([]byte, error, bool) {
	if REPO == nil {
		return []byte(""), errors.New("Internal Server Error: The repo was not specified"), false
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", REPO.User, REPO.RepoName, SHA, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return []byte(""), err, false
	}

	res, err := client.Do(req)

	if err != nil {
		return []byte(""), err, false
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return []byte(""), err, false
	}

	return body, nil, res.StatusCode == 404
}
