package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}

func FetchIssues(repo, state string, maxPages int) ([]Issue, error) {
    var allIssues []Issue

	for page := 1; page <= maxPages; page++ {
        url := fmt.Sprintf("https://api.github.com/repos/%s/issues?state=%s&page=%d", repo, state, page)
        resp, err := http.Get(url)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
        }

        var issues []Issue
        if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
            return nil, err
        }

        if len(issues) == 0 {
            break // no more issues
        }

        allIssues = append(allIssues, issues...)
    }

	// This function would contain the logic to fetch issues from GitHub API
	// For now, let's return a dummy list of issues
	// url := fmt.Sprintf("https://api.github.com/repos/%s/issues?state=%s", repo, state)
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return nil, err
	// }

	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("failed to fetch issues: %s", resp.Status)
	// }

	// var issues []Issue
	// err = json.NewDecoder(resp.Body).Decode(&issues)
	// if err != nil {
	// 	return nil, err
	// }
	// Check if the issues are empty
	// If empty, return an error
	// This is just a placeholder check, you can modify it as per your requirements
	if len(allIssues) == 0 {
		return nil, fmt.Errorf("no issues found for repository %s", repo)
	}
	return allIssues, nil
}
