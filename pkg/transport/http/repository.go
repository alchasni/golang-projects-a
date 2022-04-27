package http

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"twatter/pkg/core/service/permission"
	"twatter/pkg/core/service/role"
	"twatter/pkg/core/service/rolepermission"
	"twatter/pkg/transport/http/middleware"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type HTTP struct {
	PermissionService     permission.UseCase
	RoleService           role.UseCase
	RolePermissionService rolepermission.UseCase

	Env    string
	Config Config

	e       *echo.Echo
	baseURL *url.URL
	paths   map[string]string
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (h *HTTP) Init() error {
	var err error

	e := echo.New()
	e.Use(
		echomiddleware.Logger(),
		echomiddleware.Recover(),
		middleware.CORS(),
		middleware.Headers())
	e.HTTPErrorHandler = ErrHandler{E: e}.Handler
	e.Static("/assets", "web/public/views/assets")

	h.baseURL, err = url.Parse(h.Config.BaseURL)
	if err != nil {
		return err
	}

	h.e = e
	h.paths = make(map[string]string)

	h.register()

	return nil
}

func (h HTTP) Serve() {
	s := &http.Server{
		Addr:         h.Config.Port,
		ReadTimeout:  time.Duration(h.Config.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(h.Config.WriteTimeoutSeconds) * time.Second,
	}
	h.e.Debug = h.Config.Debug

	// Start server
	quit := make(chan os.Signal)
	go func() {
		if err := h.e.StartServer(s); err != nil {
			h.e.Logger.Error(err)
			h.e.Logger.Info("shutting down the server...")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 10 seconds.
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := h.e.Shutdown(ctx); err != nil {
		h.e.Logger.Fatal(err)
	}
}
