package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type pipeline string

func (g pipeline) CreatePipeline() {
	fmt.Println("Creating Buildkite Pipeline")
}

// Pipeline exported as symbol
var Pipeline pipeline

func oldcode(pipelineName string) {
	println("Pipeline Name: " + pipelineName)

	// Init vars and consts
	const uri string = "https://api.buildkite.com"
	const homepage string = "http://lastknight.co.uk"

	// First Check Env Variables and use these
	var BKOrg = os.Getenv("XO_BUILDKITE_ORG")
	var BKToken = os.Getenv("XO_BUILDKITE_TOKEN")
	var GHOrg = os.Getenv("XO_GITHUB_ORG")
	var GHRepoName = "git@github.com:" + GHOrg + "/" + pipelineName + ".git"
	println("GHRPEO: " + GHRepoName)

	type Step struct {
		Type    string `json:"type"`
		Name    string `json:"name"`
		Command string `json:"command"`
	}

	type Pipeline struct {
		Name       string `json:"name"`
		Repository string `json:"repository"`
		Steps      []Step `json:"steps"`
	}

	step := Step{"script", "pipeline:", "buildkite-agent pipeline upload"}
	data := Pipeline{Name: pipelineName, Repository: GHRepoName}
	data.Steps = append(data.Steps, step)

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)
	apiTarget := uri + "/v2/organizations/" + BKOrg + "/pipelines"
	println("api:" + apiTarget)
	req, err := http.NewRequest("POST", apiTarget, body)
	if err != nil {
		// handle err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+BKToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {

	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//
		}
		bodyString := string(bodyBytes)
		println(bodyString)
	}

	defer resp.Body.Close()
}
