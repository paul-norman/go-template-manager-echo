# `templateManager` Echo Integration

This is the package for the [`templateManager`](https://github.com/paul-norman/go-template-manager) renderer integration for the [Echo](https://echo.labstack.com) framework.

For all options, please see the main repository.

## Basic Usage

```go
package main

import (
	"net/http"

	TM "github.com/paul-norman/go-template-manager-echo"
	"github.com/labstack/echo/v4"
)

func main() {
	renderer := TM.Init("templates", ".html")
	renderer.ExcludeDirectories([]string{"layouts", "partials"}).
			Reload(true).
			Debug(true).
			Parse()
	
	e := echo.New()
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", TM.Params{
			"Title": "Home",
		})
	})

	e.GET("/test", func(c echo.Context) error {
		return c.Render(http.StatusOK, "test.html", TM.Params{
			"Title": "Test",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```