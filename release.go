package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var _cache = make(map[string]cache)

// Release structure
type Release struct {
	Assets []struct {
		Name        string `json:"name"`
		Link        string `json:"browser_download_url"`
		ContentType string `json:"content_type"`
	} `json:"assets"`
}

type cache struct {
	rel     Release
	created time.Time
}

func _getRelease(url string) (release Release, err error) {
	for key, hit := range _cache {
		if key == url {
			if time.Since(hit.created).Minutes() > time.Duration.Minutes(5) {
				// if cache is 5 min old clear it
				delete(_cache, key)
			} else {
				return hit.rel, nil
			}
		}
	}
	r, err := http.Get(url)
	if err != nil {
		return
	}
	json.NewDecoder(r.Body).Decode(&release)
	_cache[url] = cache{
		rel:     release,
		created: time.Now(),
	}
	return
}

func getRelease(repo string, tag string) (Release, error) {
	return _getRelease(fmt.Sprintf("https://api.github.com/repos/%s/releases/tags/%s", repo, tag))
}

func getLatestRelease(repo string) (Release, error) {
	return _getRelease(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo))
}
