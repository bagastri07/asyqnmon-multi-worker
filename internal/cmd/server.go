package cmd

import (
	"html/template"
	"io"

	"github.com/bagastri07/asyqnmon-multi-worker/internal/config"
	"github.com/bagastri07/asyqnmon-multi-worker/internal/handler"
	"github.com/bagastri07/asyqnmon-multi-worker/internal/provider"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func runServer(cmd *cobra.Command, args []string) {
	e := echo.New()

	// use middleware
	e.Use(middleware.Logger())

	// Parse the HTML templates
	templates := template.Must(template.ParseGlob("view/*.html"))

	// Set up the custom template renderer
	renderer := &Template{templates: templates}
	e.Renderer = renderer

	// Serve static files (CSS, JS, etc.)
	e.Static("/static", "static")

	// provider
	aysnqProvider := provider.NewAsyqnmonProvider()

	// handler
	handler := handler.NewHandler(aysnqProvider)

	// Define a route for the home page
	e.GET("/", handler.RootHandler)
	e.GET("/user-worker/*", handler.UserWorkerHandler)
	e.GET("company-worker/*", handler.CompanyWorkerHandler)
	e.GET("notification-worker/*", handler.NotificationWorkerHandler)

	// Start the server
	e.Logger.Fatal(e.Start(":" + config.GetPort()))
}

func init() {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start the server",
		Run:   runServer,
	}

	rootCmd.AddCommand(serverCmd)
}
