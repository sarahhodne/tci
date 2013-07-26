package travis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const TRAVIS_API_URL = "https://api.travis-ci.org"

type Repository struct {
	ID          int
	LastBuildID int
}

type Build struct {
	ID            int
	Number        string
	CommitSubject string
	State         string
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

func (c TravisClient) GetRepository(slug string) (Repository, error) {
	resp, err := NewRequest(c, fmt.Sprintf("repos/%s", slug), "")
	if err != nil {
		return Repository{}, err
	}

	if resp["repo"] == nil {
		return Repository{}, errors.New("GetRepository: Could not find repository")
	}

	repoInfo := resp["repo"].(map[string]interface{})
	repo := Repository{
		ID:          int(repoInfo["id"].(float64)),
		LastBuildID: int(repoInfo["last_build_id"].(float64)),
	}

	return repo, nil
}

func (c TravisClient) GetBuild(id int) (Build, error) {
	resp, err := NewRequest(c, fmt.Sprintf("builds/%d", id), "")
	if err != nil {
		return Build{}, err
	}

	if resp["build"] == nil {
		return Build{}, errors.New("GetBuild: Could not find build")
	}

	buildInfo := resp["build"].(map[string]interface{})
	build := Build{
		ID:            int(buildInfo["id"].(float64)),
		Number:        buildInfo["number"].(string),
		CommitSubject: resp["commit"].(map[string]interface{})["message"].(string),
		State:         buildInfo["state"].(string),
	}

	return build, nil
}

func NewRequest(c TravisClient, path string, params string) (map[string]interface{}, error) {
	client := c.client
	url := fmt.Sprintf("%s/%s?%s", c.BaseURL, path, params)

	var decodedResponse map[string]interface{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return decodedResponse, err
	}

	req.Header.Set("Accept", "application/json; version=2")

	resp, err := client.Do(req)
	if err != nil {
		return decodedResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		return decodedResponse, err
	}

	err = json.Unmarshal(body, &decodedResponse)

	// Check for bad JSON
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to decode JSON response (HTTP %v): %s", resp.StatusCode, body))
		return decodedResponse, err
	}

	return decodedResponse, nil
}
