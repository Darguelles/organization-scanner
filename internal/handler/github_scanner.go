package handler

import (
	"github.com/sirupsen/logrus"
	"organization-scanner/internal/github"
	"organization-scanner/internal/repository"
	"organization-scanner/internal/scanner"
)

type GitHubRepositoryScanner interface {
	ScanRepositoriesFromOrganization(org *string)
	ScanRepository(repositoryURL *string)
}

type gitHubRepositoryScanner struct {}

func NewGitHubRepositoryScanner() GitHubRepositoryScanner {
	return gitHubRepositoryScanner{}
}

func (g gitHubRepositoryScanner) ScanRepositoriesFromOrganization(org *string) {
	err := gitHubService().ScanRepositoriesFromOrganization(org)
	if err != nil {
		logrus.Error("Error scanning repositories", err)
	}
}

func (g gitHubRepositoryScanner) ScanRepository(repositoryURL *string) {
	err := gitHubService().ScanRepository(repositoryURL)
	if err != nil {
		logrus.Error("Error scanning repository", err)
	}
}

func gitHubService() repository.Service{
	git := github.NewGitHubService()
	scan := scanner.NewScanService()
	service := repository.NewRepositoryService(git, scan)
	return service
}
