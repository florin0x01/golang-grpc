package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetGithubResponse generates response to a GithubRequest
func (s *Server) GetGithubResponse(ctx context.Context, in *GithubRequest) (*GithubResponse, error) {
	inputQuery := fmt.Sprintf("query { viewer { repositories(first: %s, affiliations: [OWNER]) { totalCount pageInfo { endCursor hasNextPage } nodes { name owner { login } } } } }", in.Message)
	requestBody, err := json.Marshal(map[string]string{
		"query": inputQuery,
	})

	if err != nil {
		return &GithubResponse{Message: err.Error()}, nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", "bearer <github_token_here>")
	resp, err := client.Do(req)
	if err != nil {
		return &GithubResponse{Message: err.Error()}, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &GithubResponse{Message: err.Error()}, nil
	}

	respBody := string(body)

	if err != nil {
		return &GithubResponse{Message: err.Error()}, nil
	}

	return &GithubResponse{Message: respBody}, nil
}
