package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Release structure
type Release struct {
	Assets []struct {
		Name        string `json:"name"`
		Link        string `json:"browser_download_url"`
		ContentType string `json:"content_type"`
	} `json:"assets"`
}

func _getRelease(url string) (release Release, err error) {
	r, err := http.Get(url)
	if err != nil {
		return
	}
	json.NewDecoder(r.Body).Decode(&release)
	return
}

func getRelease(repo string, tag string) (Release, error) {
	return _getRelease(fmt.Sprintf("https://api.github.com/repos/%s/releases/tags/%s", repo, tag))
}

func getLatestRelease(repo string) (Release, error) {
	return _getRelease(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo))
}
