package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type pipeline string

const uri string = "https://api.buildkite.com"

func (g pipeline) CreatePipeline(name string) {
	_CreatePipeline(name)
}

// Pipeline exported as symbol
var Pipeline pipeline
var pipelineOrg = os.Getenv("XO_PIPELINE_ORG")
var pipelineToken = os.Getenv("XO_PIPELINE_TOKEN")
var repoOrg = os.Getenv("XO_REPO_ORG")

// Plugin Implementation

func _CreatePipeline(name string) {

	// Init vars and consts

	var repoName = "git@github.com:" + repoOrg + "/" + name + ".git"

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
	data := Pipeline{Name: name, Repository: repoName}
	data.Steps = append(data.Steps, step)

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body := bytes.NewReader(payloadBytes)
	apiTarget := uri + "/v2/organizations/" + pipelineOrg + "/pipelines"

	req, err := http.NewRequest("POST", apiTarget, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+pipelineToken)

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
