package controller

import (
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	if _, err := os.Stat("upload"); os.IsNotExist(err) {
		err = os.Mkdir("upload", fs.ModePerm)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(filepath.Join("upload", filepath.Base(file.Filename)))
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
