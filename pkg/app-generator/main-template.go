package appgenerator

var mainTemplate = `
package main

import (
	"log"
	{{.AppName}}Repo "{{.Repo}}/pkg/repo/{{.AppName}}"
	{{.AppName}}Service "{{.Repo}}/pkg/service/{{.AppName}}"
)

type appConfig struct {
	appName string
	upstreamURL string
}

func main(){
	repo,err := provide{{.AppName}}Repo()
	if err != nil {
		log.Panicln("couldn't initialize {{.AppName}}Repo",err)
	}
	service := {{.AppName}}Service.New(repo)
	return
}

func provide{{.AppName}}Repo() (repo *{{.AppName}}Repo,err error) {
	// construct and pass repo configuration here
	repo =  {{.AppName}}Repo.New(appConfig.upstreamURL)
	err = repo.Initialize()

	return
}
`
