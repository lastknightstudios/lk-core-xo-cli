package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {

	//Create GitHub Project

	// Set expected Vars

	var xoGitHubOrg string
	var xoGitHubToken string

	// Handling Authentication

	// First Check Env Variables and use these
	xoGitHubOrg = os.Getenv("XO_GITHUB_ORG")
	xoGitHubToken = os.Getenv("XO_GITHUB_TOKEN")

	print(xoGitHubOrg)
	print(xoGitHubToken)

	type Payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
		Private     bool   `json:"private"`
		HasIssues   bool   `json:"has_issues"`
		HasProjects bool   `json:"has_projects"`
		HasWiki     bool   `json:"has_wiki"`
	}

	data := Payload{"lk-core-go-test-2", "A Test Repository", "http://lastknight.co.uk", true, false, false, false}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.github.com/orgs/"+xoGitHubOrg+"/repos", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+xoGitHubToken)

	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}
