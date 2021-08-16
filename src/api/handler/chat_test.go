package handler

import (
	"droneia-go/src/api/db"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	chatJSON = `[{"name":"Jon Snow"}]`
)

func TestGetChat(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/chats")

	db, _ := db.NewMongoDB(os.Getenv("MONGO_DB_URL"))
	// TODO: change set DB from variables to argument for testing
	h := &Handler{
		DB: db,
	}

	if assert.NoError(t, h.GetChat(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, chatJSON, rec.Body.String())
	}
}
