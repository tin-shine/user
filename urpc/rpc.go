package urpc

import (
	"context"
	"log"

	"github.com/game_marketing/user/proto/user"
)

type Server struct {
	user.UnimplementedInfoServer
}

func (s *Server) GetInfo(c context.Context, in *user.InfoRequest) (*user.InfoResponse, error) {
	log.Println("get.user.info")
	return &user.InfoResponse{
		Uid:  123,
		Name: "tinshine",
		Sex:  "male",
	}, nil
}
