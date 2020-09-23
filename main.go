package main

import (
	"organization-scanner/internal/github"
	"organization-scanner/internal/repository"
)

func main() {
	git := github.NewGitHubService()
	service := repository.NewRepositoryService(git)
	service.ListRepositories()
}
