// Package services coordinates application business services.
package services

import "github.com/bhcoder23/gin-layout/internal/repos"

type Services func()

var services = []Services{}

func InitServices() {
	repos.InitRepos()
	for _, service := range services {
		service()
	}
}

func RegisterServices(r ...Services) {
	services = append(services, r...)
}
