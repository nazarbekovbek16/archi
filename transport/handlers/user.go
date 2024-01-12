package handlers

import (
	"archi/config"
	"archi/service"
	"archi/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserHandler struct {
	storage     *storage.Storage
	userManager *service.Service
	config      *config.Config
	log         *zap.Logger
}

func NewUserHandler(logger *zap.Logger, config *config.Config, storage *storage.Storage, userManager *service.Service) *UserHandler {
	return &UserHandler{log: logger, config: config, storage: storage, userManager: userManager}
}

func (h *UserHandler) SignIn(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"IsError": false,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "bek" && password == "1" {
		cookie := new(http.Cookie)

		cookie.Name = "token"
		cookie.Value = "token"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		h.log.Info("Token created")
		return c.Redirect(http.StatusFound, "/")
	} else {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"IsError": true,
		})
	}

}

func (h *UserHandler) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Expires = time.Now().Add(6 * time.Hour)
	c.SetCookie(cookie)

	referringURL := c.Request().Header.Get("Referer")

	return c.Redirect(http.StatusFound, referringURL)
}

func (h *Handlers) MainPage(c echo.Context) error {
	token, _ := ReadCookie(c)
	isAuth := false
	if token == "token" {
		isAuth = true
	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"IsAuthenticated": isAuth,
	})
}
func (h *Handlers) Map(c echo.Context) error {
	token, _ := ReadCookie(c)
	isAuth := false
	if token == "token" {
		isAuth = true
	}
	return c.Render(http.StatusOK, "map.html", map[string]interface{}{
		"IsAuthenticated": isAuth,
	})
}
func (h *Handlers) KMD(c echo.Context) error {
	token, _ := ReadCookie(c)
	isAuth := false
	if token == "token" {
		isAuth = true
	}
	return c.Render(http.StatusOK, "kmd.html", map[string]interface{}{
		"IsAuthenticated": isAuth,
	})
}
func (h *Handlers) Service(c echo.Context) error {
	token, _ := ReadCookie(c)
	isAuth := false
	if token == "token" {
		isAuth = true
	}
	return c.Render(http.StatusOK, "service.html", map[string]interface{}{
		"IsAuthenticated": isAuth,
	})
}
func (h *Handlers) ArchModel(c echo.Context) error {
	token, _ := ReadCookie(c)
	isAuth := false
	if token == "token" {
		isAuth = true
	}
	return c.Render(http.StatusOK, "arch_model.html", map[string]interface{}{
		"IsAuthenticated": isAuth,
	})
}
func (h *Handlers) KMDLegend(c echo.Context) error {
	return c.Render(http.StatusOK, "kmd_legend.html", nil)
}
func (h *Handlers) MapLegend(c echo.Context) error {
	return c.Render(http.StatusOK, "map_legend.html", nil)
}
func (h *Handlers) ServiceLegend(c echo.Context) error {
	return c.Render(http.StatusOK, "service_legend.html", nil)
}
func (h *Handlers) ModelKMD(c echo.Context) error {
	return c.Render(http.StatusOK, "model.html", nil)
}
func (h *Handlers) ModelMap(c echo.Context) error {
	return c.Render(http.StatusOK, "modelMap.html", nil)
}

func (h *Handlers) Test(c echo.Context) error {
	id := c.Param("id")
	elements, _ := h.User.userManager.Element.GetByID(c.Request().Context(), id)

	//fmt.Println(elements)
	return c.Render(http.StatusOK, "attribute.html", elements)
}
