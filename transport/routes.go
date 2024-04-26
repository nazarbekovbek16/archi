package transport

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (s *Server) routes() {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	//s.HTTP.Static("/", "C:/Users/kzhol/GolandProjects/archi/transport/templates")

	// Middleware
	//s.HTTP.Use(middleware.Logger())
	//s.HTTP.Use(middleware.Recover())

	s.HTTP.Static("/", "templates")
	////s.HTTP.Static("/", "C:/Users/kzhol/GolandProjects/archi/transport/templates/id-89979e64baae4e45ac5fc7f3438e8d8e/elements")

	s.HTTP.Static("/assets", "templates/assets")
	s.HTTP.Static("/lib", "templates/lib")
	s.HTTP.Static("/css", "templates/css")
	s.HTTP.Static("/js", "templates/js")

	s.HTTP.Renderer = renderer

	s.HTTP.GET("/login", s.h.User.SignIn)
	s.HTTP.POST("/login", s.h.User.Login)
	s.HTTP.GET("/logout", s.h.User.Logout)

	//s.HTTP.GET("/tt", s.h.Test)

	s.HTTP.GET("/element/:id", s.h.Test)

	s.HTTP.GET("/", s.h.MainPage)
	s.HTTP.GET("/map", s.h.Map)
	s.HTTP.GET("/kmd", s.h.KMD)
	s.HTTP.GET("/service", s.h.Service)
	s.HTTP.GET("/arch_model", s.h.ArchModel)

	s.HTTP.GET("/kmd_legend", s.h.KMDLegend)
	s.HTTP.GET("/map_legend", s.h.MapLegend)
	s.HTTP.GET("/service_legend", s.h.ServiceLegend)

	s.HTTP.GET("/modelKMD", s.h.ModelKMD)
	s.HTTP.GET("/modelMap", s.h.ModelMap)
}
