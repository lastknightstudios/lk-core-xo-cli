package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type repository string

func (g repository) CreateRepository(name string) {
	fmt.Println("Creating Repository", name)
	_CreateRepository(name)
}

func (g repository) CreateWebhook() {
	fmt.Println("Creating WebHook")
}

// Repository exported as symbol
var Repository repository

func _CreateRepository(name string) {

	// Init vars and consts
	const uri string = "https://api.github.com"
	var GHOrg string
	var GHToken string

	// First Check Env Variables and use these
	GHOrg = os.Getenv("XO_GITHUB_ORG")
	GHToken = os.Getenv("XO_GITHUB_TOKEN")

	var GHRepoName = name
	var GHRepoPrivate = true
	var GHRepoHasIssues = false
	var GHRepoHasProjects = false
	var GHRepoHasWiki = false
	//var GHRepoTeamID = "lk-core-developers"

	type _Repository struct {
		Name        string `json:"name"`
		Private     bool   `json:"private"`
		HasIssues   bool   `json:"has_issues"`
		HasProjects bool   `json:"has_projects"`
		HasWiki     bool   `json:"has_wiki"`
	}

	data := _Repository{GHRepoName, GHRepoPrivate, GHRepoHasIssues, GHRepoHasProjects, GHRepoHasWiki}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", uri+"/orgs/"+GHOrg+"/repos", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GHToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		// handle err
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
