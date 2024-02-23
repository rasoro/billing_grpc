package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/google/uuid"
	pb "github.com/nyaruka/courier/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultContactUUID = "b62d6073-b14f-48aa-a75d-7968f7c8badd"
	defaultChannelUUID = "25976982-a081-4b2f-9c41-3adb257db429"
)

var (
	addr         = flag.String("addr", "localhost:50051", "the address to connect to")
	contactUUID  = flag.String("contact_uuid", defaultContactUUID, "contact uuid")
	channelUUIUD = flag.String("channel_uuid", defaultChannelUUID, "contact uuid")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFlowsMesssageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateMessage(ctx, &pb.CreateFlowsMessageRequest{
		ContactUuid: uuid.New().String(),
		ChannelUuid: uuid.New().String(),
		MessageDate: time.Now().Format("2006-01-02 15:04:05"),
		MessageUuid: uuid.New().String(),
	})
	if err != nil {
		log.Fatalf("could not create message: %v", err)
	}
	log.Printf("Message: %s", r.GetMessage())

}
