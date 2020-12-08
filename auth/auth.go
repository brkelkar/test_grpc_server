package auth

import (
	context "context"
	"log"
)

type Server struct {
}

//Validate just validates
func (s *Server) Validate(ctx context.Context, in *Request) (*Responce, error) {
	log.Printf("Receive message body from client: %s", in.AuthToken)
	return &Responce{Valid: true}, nil
}

func (s *Server) mustEmbedUnimplementedAuthValidationServiceServer() {}
