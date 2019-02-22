package main

import (
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
		return _dl(c, rel)
	})
	e.GET("/dl/:tag/:file", func(c echo.Context) error {
		rel, err := getRelease(repo, c.Param("tag"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return _dl(c, rel)
	})
	e.Start(addr)
}

func _dl(c echo.Context, rel Release) error {
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
}
