package handler

import (
	"github.com/abhisheksinghvi09/go-boilerplate/internal/server"
	"github.com/abhisheksinghvi09/go-boilerplate/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
