package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func proxy(repo string, addr string) {
	e := echo.New()

	e.GET("/dl/:file", func(c echo.Context) error {
		rel, err := getLatestRelease(repo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		for _, asset := range rel.Assets {
			if asset.Name == c.Param("file") {
				res, err := http.Get(asset.Link)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}
				defer res.Body.Close()
				return c.Stream(http.StatusOK, asset.ContentType, res.Body)
			}
		}
		return c.JSON(http.StatusNotFound, "file not found")
	})

	e.GET("/dl/:tag/:file", func(c echo.Context) error {
		r, err := http.Get(fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", repo, c.Param("tag"), c.Param("file")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer r.Body.Close()
		contentType := r.Header.Get("Content-type")
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		return c.Stream(r.StatusCode, contentType, r.Body)
	})

	e.Start(addr)
}
