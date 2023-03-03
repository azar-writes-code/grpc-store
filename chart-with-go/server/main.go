package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "server/charts"
)

// type server struct {
// 	pb.UnimplementedChartDataServiceServer
// }

// func (s *server) GetLineChartData(req *pb.ChartRequest, stream pb.ChartDataService_GetLineChartDataServer) error {
// 	for{
// 		if err := stream.Send(&pb.ChartResponse{
// 			X : time.Now().Unix(),
// 			Y : int64(rand.Intn(150 - 30 + 1) + 30),
// 		}); err != nil {
// 			return err
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	return nil
// }

// func (s *server) GetBarChartData(req *pb.ChartRequest, stream pb.ChartDataService_GetBarChartDataServer) error {
// 	for{
// 		if err := stream.Send(&pb.ChartResponse{
// 			X : time.Now().Unix(),
// 			Y : int64(rand.Intn(150 - 30 + 1) + 30),
// 		}); err != nil {
// 			return err
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	return nil
// }

var (
	port int = 8080
)

func main() {

	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChartDataServiceServer(s, &helpers.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("Starting server in port :%d\n", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
