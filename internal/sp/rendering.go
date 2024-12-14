package sp

import (
	"net/http"
)

type Renderer interface {
	Render(w http.ResponseWriter, name []string, data interface{}) error
}

type RenderingService struct {
	Renderer Renderer
}

func NewRenderingService(r Renderer) RenderingService {
	return RenderingService{
		Renderer: r,
	}
}
