package controller

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/labstack/echo/v4"
)

// Read excel data
func readExcelFile(src io.Reader) error {
	f, err := excelize.OpenReader(src)
	if err != nil {
		return err
	}
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}
	for _, row := range rows[1:] {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	return nil
}

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
	readExcelFile(src)

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
