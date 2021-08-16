package handler

import (
	"droneia-go/src/api/db"
	"droneia-go/src/api/model"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// TODO: fixing return response withouht new line
	chatsJSON = `[{"name":"Test","status":"ING"}]
`
	chatJSON = `{"_id":"1","name":"Test","status":"ING"}
`
)

type TestChatService struct {
	DB *mongo.Client
}

func (c *TestChatService) FindAll(options *options.FindOptions) ([]*model.Chat, error) {
	var results []*model.Chat
	results = append(results, &model.Chat{
		Name:   "Test",
		Status: "ING",
	})
	return results, nil
}

func (c *TestChatService) InsertOne(data *model.Chat) (*model.Chat, error) {
	data.Id = "1"
	return data, nil
}

func TestGetChat(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/chats")

	db, _ := db.NewMongoDB(os.Getenv("MONGO_DB_URL"))
	// TODO: change set DB from variables to argument for testing
	h := &Handler{
		DB: db,
		ChatService: &TestChatService{
			DB: db,
		},
	}
	h.InitService()

	if assert.NoError(t, h.GetChats(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, chatsJSON, rec.Body.String())
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func TestCreateChat(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(chatJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/chats")

	db, _ := db.NewMongoDB(os.Getenv("MONGO_DB_URL"))
	// TODO: change set DB from variables to argument for testing
	h := &Handler{
		DB: db,
		ChatService: &TestChatService{
			DB: db,
		},
	}
	h.InitService()

	if assert.NoError(t, h.CreateChat(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, chatJSON, rec.Body.String())
	}
}
