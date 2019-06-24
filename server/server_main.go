package server

import (
	"fmt"
	"log"
	"net"
	"os"
	qpb "perScoreCal/perScoreProto/question"
	upb "perScoreCal/perScoreProto/user"

	"google.golang.org/grpc"
)

// StartServer ...
func StartServer() {
	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	upb.RegisterUserServer(s, &UserServer{})
	qpb.RegisterQuestionServer(s, &QuestionServer{})
	fmt.Println("perScoreCal server started on", os.Getenv("PORT"))
	s.Serve(lis)
}
