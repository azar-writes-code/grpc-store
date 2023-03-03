package helpers

import (
	"fmt"
	"time"
	"math/rand"
	pb "server/charts"
)

type Server struct {
	pb.UnimplementedChartDataServiceServer
} 

func GenerateData()  {
	fmt.Println("hello World")
}

func (s *Server) GetLineChartData(req *pb.ChartRequest, stream pb.ChartDataService_GetLineChartDataServer) error {
	for{
		if err := stream.Send(&pb.ChartResponse{
			X : time.Now().UnixMilli(),
			Y : int64(rand.Intn(150 - 30 + 1) + 30),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 4)
	}
	return nil
} 

func (s *Server) GetBarChartData(req *pb.ChartRequest, stream pb.ChartDataService_GetBarChartDataServer) error {
	for{
		if err := stream.Send(&pb.ChartResponse{
			X : time.Now().UnixMilli(),
			Y : int64(rand.Intn(150 - 30 + 1) + 30),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 4)
	}
	return nil
}