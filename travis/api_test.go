package travis

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	client *TravisClient
	server *httptest.Server
)

func testRequest(t *testing.T, r *http.Request, method, path string) {
	if r.Method != method {
		t.Errorf("Expected method %v, got %v", method, r.Method)
	}
	if r.URL.Path != path {
		t.Errorf("Expected path %v, got %v", path, r.URL.Path)
	}
}

func stubRequest(t *testing.T, method, path, body string) {
	stubRequestWithStatus(t, method, path, body, http.StatusOK)
}

func stubRequestWithStatus(t *testing.T, method, path, body string, statusCode int) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testRequest(t, r, method, path)
		w.WriteHeader(statusCode)
		fmt.Fprint(w, body)
	}))
}

func teardown() {
	server.Close()
}

func TestGetRepository(t *testing.T) {
	stubRequest(t, "GET", "/repos/foo/bar", `{"repo":{"id":123,"last_build_id":234}}`)
	defer teardown()

	client := TravisClient{client: http.DefaultClient, BaseURL: server.URL}
	repo, err := client.GetRepository("foo/bar")
	if err != nil {
		t.Errorf("client.GetRepository errored: %v", err)
	}

	want := Repository{
		ID:          123,
		LastBuildID: 234,
	}

	if !reflect.DeepEqual(repo, want) {
		t.Errorf("client.GetRepository = %v, want %v", repo, want)
	}
}

func TestGetRepository_NotFound(t *testing.T) {
	stubRequestWithStatus(t, "GET", "/repos/foo/bar", `{}`, http.StatusNotFound)
	defer teardown()

	client := TravisClient{client: http.DefaultClient, BaseURL: server.URL}
	repo, _ := client.GetRepository("foo/bar")

	if !reflect.DeepEqual(repo, Repository{}) {
		t.Errorf("client.GetRepository: expected empty repo, got %v", repo)
	}
}

func TestGetBuild(t *testing.T) {
	stubRequest(t, "GET", "/builds/234", `{"build":{"id":234,"number":"1","state":"passed"},"commit":{"message":"Hello, world"}}`)
	defer teardown()

	client := TravisClient{client: http.DefaultClient, BaseURL: server.URL}
	build, err := client.GetBuild(234)
	if err != nil {
		t.Errorf("client.GetBuild errored: %v", err)
	}

	want := Build{
		ID:     234,
		Number: "1",
		State:  "passed",
	}

	if !reflect.DeepEqual(build, want) {
		t.Errorf("client.GetBuild = %v, want %v", build, want)
	}
}
