package main
import (
	// "fmt"
	"math/rand"
	"log"
	"net"
	"time"
	"strconv"
	"google.golang.org/grpc"
	pb "server/charts"
)

type chart struct{
	pb.UnimplementedChartDataServiceServer
}

func (*chart) GetLineChartData(req *pb.LineChartRequest,
	stream pb.ChartDataService_GetLineChartDataServer) error {
		for i := int64(0); i < 10; i++ {
			if err := stream.Send(&pb.LineChartResponse{
				X: i,
 				Y: rand.Int63n(100),
			}); err != nil {
				return err
			}
			time.Sleep(time.Second)
		}

	return nil
}

func (*chart) GetProgressChartData(req *pb.ProgressChartRequest, stream pb.ChartDataService_GetProgressChartDataServer) error {
	for i:= int64(0); i< 30 ; i++ {
		if err := stream.Send(&pb.ProgressChartResponse{
			Title: "Progress" + strconv.Itoa(int(i)),
			Value : rand.Int63n(100),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Error while listening: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChartDataServiceServer(s, &chart{})
	log.Println("start server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}