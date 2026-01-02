package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/keshuook/keshuook-web-archive/internal/githubapi"
	"github.com/keshuook/keshuook-web-archive/internal/server"
)

var REPO githubapi.Repo = githubapi.Repo{User: "keshuook", RepoName: "keshuook.github.io"}

func main() {
	// Constant and env
	err := godotenv.Load(".env")

	// Parse REPO constant to server
	server.SetRepo(&REPO)

	if err != nil {
		fmt.Println(err)
	}

	// Get list of commits
	githubapi.GetCommits(REPO)

	var port int

	if os.Getenv("PORT") == "" {
		port = 3000 // Fallback for local dev
	} else {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}

	// Server
	server.StartServer(port)
}
