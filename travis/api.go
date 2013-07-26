package travis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const TRAVIS_API_URL = "https://api.travis-ci.org"

type Repository struct {
	ID          int `json:"id"`
	LastBuildID int `json:"last_build_id"`
}

type Build struct {
	ID          int    `json:"id"`
	Number      string `json:"number"`
	State       string `json:"state"`
	PullRequest bool   `json:"pull_request"`
	Duration    int    `json:"duration"`
	StartedAt   string `json:"started_at"`
	FinishedAt  string `json:"finished_at"`
}

type Commit struct {
	Message    string `json:"message"`
	Branch     string `json:"branch"`
	CompareURL string `json:"compare_url"`
}

type BuildResponse struct {
	Build  Build  `json:"build"`
	Commit Commit `json:"commit"`
}

type RepositoryResponse struct {
	Repository Repository `json:"repo"`
}

type TravisClient struct {
	client *http.Client

	BaseURL string
}

func NewClient() *TravisClient {
	c := &TravisClient{
		client:  http.DefaultClient,
		BaseURL: TRAVIS_API_URL,
	}
	return c
}

func (c TravisClient) GetRepository(slug string) (RepositoryResponse, error) {
	body, err := NewRequest(c, fmt.Sprintf("repos/%s", slug), "")
	if err != nil {
		return RepositoryResponse{}, err
	}

	var repo RepositoryResponse
	err = json.Unmarshal(body, &repo)

	return repo, err
}

func (c TravisClient) GetBuild(id int) (BuildResponse, error) {
	body, err := NewRequest(c, fmt.Sprintf("builds/%d", id), "")
	if err != nil {
		return BuildResponse{}, err
	}

	var build BuildResponse
	err = json.Unmarshal(body, &build)

	return build, err
}

func NewRequest(c TravisClient, path string, params string) ([]byte, error) {
	client := c.client
	url := fmt.Sprintf("%s/%s?%s", c.BaseURL, path, params)

	var body []byte

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}

	req.Header.Set("Accept", "application/json; version=2")

	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}

	body, err = ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	return body, err
}
