package sp

import (
	"net/http"
)

type Renderer interface {
	Render(w http.ResponseWriter, r *http.Request, data interface{})
}

type RenderingService struct {
	Renderer Renderer
}

func NewRenderingService(r Renderer) RenderingService {
	return RenderingService{
		Renderer: r,
	}
}
