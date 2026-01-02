package server

import (
	"net/http"

	"github.com/keshuook/keshuook-web-archive/internal/githubapi"
)

func APIFirstCommit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(githubapi.CommitsList[len(githubapi.CommitsList)-1].Time.Format("2006-01-02")))
}

func APILastCommit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(githubapi.CommitsList[0].Time.Format("2006-01-02")))
}
