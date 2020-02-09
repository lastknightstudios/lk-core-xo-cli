package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func github(repoName string) {

	println("Repo Name: " + repoName)

	// Init vars and consts
	const uri string = "https://api.github.com"
	const homepage string = "http://lastknight.co.uk"
	var GHOrg string
	var GHToken string

	// First Check Env Variables and use these
	GHOrg = os.Getenv("XO_GITHUB_ORG")
	GHToken = os.Getenv("XO_GITHUB_TOKEN")

	var GHRepoName = repoName
	var GHEHomepage = homepage
	var GHRepoPrivate = true
	var GHRepoHasIssues = false
	var GHRepoHasProjects = false
	var GHRepoHasWiki = false
	//var GHRepoTeamID = "lk-core-developers"

	type Repository struct {
		Name        string `json:"name"`
		Homepage    string `json:"homepage"`
		Private     bool   `json:"private"`
		HasIssues   bool   `json:"has_issues"`
		HasProjects bool   `json:"has_projects"`
		HasWiki     bool   `json:"has_wiki"`
	}

	data := Repository{GHRepoName, GHEHomepage, GHRepoPrivate, GHRepoHasIssues, GHRepoHasProjects, GHRepoHasWiki}

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
	}
	defer resp.Body.Close()

}

func webhook(webhook string) {
	println("Webhook URI: " + webhook)

	// Init vars and consts
	const uri string = "https://api.github.com"
	const homepage string = "http://lastknight.co.uk"
	var GHOrg string
	var GHToken string

	// First Check Env Variables and use these
	GHOrg = os.Getenv("XO_GITHUB_ORG")
	GHToken = os.Getenv("XO_GITHUB_TOKEN")

	type WebHook struct {
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			URL         string `json:"url"`
			ContentType string `json:"content_type"`
			InsecureSsl string `json:"insecure_ssl"`
		} `json:"config"`
	}

}
