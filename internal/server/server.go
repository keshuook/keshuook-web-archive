package server

import (
	"fmt"
	"net/http"

	"github.com/keshuook/keshuook-web-archive/internal/middleware"
)

func StartServer(port int) {
	mux := http.NewServeMux()
	fileRouter := http.FileServer(http.Dir("./web"))

	HandleArchive(mux)
	mux.HandleFunc("/api/firstcommit/", APIFirstCommit)
	mux.HandleFunc("/api/lastcommit/", APILastCommit)
	mux.Handle("/", fileRouter)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), middleware.Logger(mux))

	if err != nil {
		fmt.Println(err.Error())
	}
}
