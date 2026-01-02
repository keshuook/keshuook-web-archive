package server

import (
	"bytes"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/keshuook/keshuook-web-archive/internal/fetch"
	"github.com/keshuook/keshuook-web-archive/internal/githubapi"
	"github.com/keshuook/keshuook-web-archive/internal/search"
)

var repo *githubapi.Repo

func SetRepo(REPO *githubapi.Repo) {
	repo = REPO
}

func injectScript(originalHTML []byte) []byte {
	// This is the script you want to inject
	snippet := []byte(`<script src="/link-fixer.js"></script>`)

	// We look for <head>. If found, we inject right after it.
	// If not found, we fallback to injecting at the very beginning.
	index := bytes.Index(originalHTML, []byte("<head>"))
	if index == -1 {
		return append(snippet, originalHTML...)
	}

	// Split at <head>, insert snippet, and join
	insertionPoint := index + len("<head>")
	newHTML := make([]byte, 0, len(originalHTML)+len(snippet))
	newHTML = append(newHTML, originalHTML[:insertionPoint]...)
	newHTML = append(newHTML, snippet...)
	newHTML = append(newHTML, originalHTML[insertionPoint:]...)

	return newHTML
}

func HandleArchive(mux *http.ServeMux) {
	mux.HandleFunc("/wayback/{year}/{month}/{day}/{reqpath...}", func(w http.ResponseWriter, r *http.Request) {
		year, e1 := strconv.ParseInt(r.PathValue("year"), 10, 32)
		month, e2 := strconv.ParseInt(r.PathValue("month"), 10, 32)
		day, e3 := strconv.ParseInt(r.PathValue("day"), 10, 32)

		if year < 0 || year > 9999 || day < 1 || day > 31 || month < 1 || month > 12 {
			w.WriteHeader(418)
		} else if e1 != nil {
			w.WriteHeader(400)
			w.Write([]byte(e1.Error()))
		} else if e2 != nil {
			w.WriteHeader(400)
			w.Write([]byte(e2.Error()))
		} else if e3 != nil {
			w.WriteHeader(400)
			w.Write([]byte(e3.Error()))
		} else {
			t := time.Date((int)(year), (time.Month)(month), (int)(day), 23, 59, 59, 0, time.Local)

			SHA, err := search.Search(githubapi.CommitsList, t)

			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
			} else {
				path := r.PathValue("reqpath")

				if strings.HasSuffix(path, "/") || path == "" {
					path = path + "index.html"
				}

				ext := filepath.Ext(path)
				mimeType := mime.TypeByExtension(ext)

				var body []byte
				var err error

				if ext != "" {
					body, err, _ = fetch.Get(SHA, path, repo)
				} else {
					var is404 bool
					body, err, is404 = fetch.Get(SHA, path+".html", repo)
					if is404 {
						http.Redirect(w, r, path+"/", http.StatusSeeOther)
						return
					}
				}

				w.Header().Set("Content-Type", mimeType)

				if ext == ".html" || ext == "" {
					body = injectScript(body)
				}

				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(err.Error()))
				} else {
					w.WriteHeader(200)
					w.Write(body)
				}
			}

		}
	})
}
