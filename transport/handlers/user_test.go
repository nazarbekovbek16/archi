package handlers

import (
	"archi/config"
	"archi/logger"
	"archi/service"
	"archi/storage"
	"archi/transport/middleware"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	createUserJSON = `{
  		"id":11,		
  		"name":"bek1",
  		"surname":"zz",
  		"email":"bek1@mail.ru",
  		"password":"12345"
	}`
)

func TestUserHandler_SignIn(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(createUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := createUserHandler(t)

	if assert.NoError(t, h.User.SignIn(c)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
		//assert.Equal(t, createUserJSON, rec.Body.String())
	}
}
func createUserHandler(t *testing.T) *Handlers {
	cfg := config.NewConfig()
	l, _ := logger.Init(cfg)
	mid := middleware.NewJWTAuth(cfg)
	repo, err := storage.NewStorage(l, context.Background(), cfg)
	if err != nil {
		t.Errorf(err.Error())
	}
	srvManager, _ := service.NewService(l, repo)
	return NewHandlers(l, cfg, repo, srvManager, mid)
}

func TestUserHandler_SignUp(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(createUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := createUserHandler(t)

	if assert.NoError(t, h.User.SignUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, createUserJSON, rec.Body.String())
	}
}
