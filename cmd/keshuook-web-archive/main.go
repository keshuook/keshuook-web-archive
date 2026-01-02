package main

import (
	"fmt"

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

	// Server
	server.StartServer(3000)
}
