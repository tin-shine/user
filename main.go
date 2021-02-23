package main

import (
	"log"
	"net"

	"github.com/game_marketing/user/proto/user"
	"github.com/game_marketing/user/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()
	register(r)
	initRPC(r)
	r.Run(":8080")
}

func register(r *gin.Engine) {
	r.GET("/ping", Ping)
}

// Ping ping
func Ping(r *gin.Context) {
	r.JSON(200, gin.H{
		"code":    0,
		"message": "ok",
	})
}

func initRPC(r *gin.Engine) {
	go func() {
		lis, err := net.Listen("tcp", ":2021")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		user.RegisterInfoServer(s, &rpc.Server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
