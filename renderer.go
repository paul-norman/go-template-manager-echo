package templateManagerEcho

import (
	"io"

	"github.com/labstack/echo/v4"
	TM "github.com/paul-norman/go-template-manager"
)

// Convenience type allowing any variables types to be passed in
type Params map[string]any

// Struct wrapper allowing override of the Render function
type Renderer struct {
	*TM.TemplateManager
}

// Creates a new Renderer instance
func Init(directory string, extension string) *Renderer {
	return &Renderer{TM.Init(directory, extension)}
}

// Renders a single template
func (r *Renderer) Render(writer io.Writer, name string, data any, c echo.Context) error {
	return r.TemplateManager.Render(name, parseData(data), writer)
}

// Converts generic variables to TM variables
func parseData(binding any) TM.Params {
	if binding == nil {
		return TM.Params{}
	}

	if old, ok := binding.(TM.Params); ok {
		return old
	}

	if old, ok := binding.(map[string]any); ok {
		return TM.Params(old)
	}

	if old, ok := binding.(Params); ok {
		return TM.Params(old)
	}

	return TM.Params{}
}