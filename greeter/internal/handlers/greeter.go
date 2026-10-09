package handlers

import (
	"context"
	"fmt"

	"greeter/internal/gen"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetHealth(ctx context.Context, request gen.GetHealthRequestObject) (gen.GetHealthResponseObject, error) {
	return gen.GetHealth200JSONResponse{Status: "ok"}, nil
}

func (s *Server) GetGreeting(ctx context.Context, request gen.GetGreetingRequestObject) (gen.GetGreetingResponseObject, error) {
	name := request.Params.Name
	if name == "" {
		name = "World"
	}
	return gen.GetGreeting200JSONResponse{Message: fmt.Sprintf("Hello, %s!", name)}, nil
}
