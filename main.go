package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuu/thermostat/bto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := bto.NewIRServiceClient(conn)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := bto.WriteRequest{
			Frequency: 38000,
			Data:      []uint32{0x00, 0x00},
		}
		res, err := client.Write(ctx, &req)
		if err != nil {
			log.Fatalf("could not write: %v", err)
		}

		c.JSON(200, gin.H{"message": "pong", "code": res.GetCode()})
	})

	r.Run(":3000")
}
