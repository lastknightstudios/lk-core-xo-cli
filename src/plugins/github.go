package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type repository string

//FI: Move these to the xo main package and pass them in instead.
var repoOrg = os.Getenv("XO_REPO_ORG")
var repoToken = os.Getenv("XO_REPO_TOKEN")

const uri string = "https://api.github.com"

func (g repository) CreateRepository(name string) {
	_CreateRepository(name)
}

func (g repository) CreateWebhook(webhook string) {
	_CreateWebhook(webhook)
}

// Repository exported as symbol
var Repository repository

// Plugin Implementation

func _CreateRepository(name string) {

	// Init vars and consts

	// First Check Env Variables and use these
	repoOrg = os.Getenv("XO_REPO_ORG")
	repoToken = os.Getenv("XO_REPO_TOKEN")

	var RepoName = name
	var RepoPrivate = true
	var RepoHasIssues = false
	var RepoHasProjects = false
	var RepoHasWiki = false
	//var RepoTeamID = "lk-core-developers"

	type _Repository struct {
		Name        string `json:"name"`
		Private     bool   `json:"private"`
		HasIssues   bool   `json:"has_issues"`
		HasProjects bool   `json:"has_projects"`
		HasWiki     bool   `json:"has_wiki"`
	}

	data := _Repository{RepoName, RepoPrivate, RepoHasIssues, RepoHasProjects, RepoHasWiki}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", uri+"/orgs/"+repoOrg+"/repos", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+repoToken)

	resp, err := http.DefaultClient.Do(req)
	scanner := bufio.NewScanner(resp.Body)

	if resp.StatusCode != 201 {
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Print(scanner.Text())
		}
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
}

func _CreateWebhook(webhook string) {
	println("Webhook URI: " + webhook)

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
