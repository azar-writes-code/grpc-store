package main

import (
	"log"
	"net"
	pb "server/sensorpb"
	"server/sensor"
	"google.golang.org/grpc"
	"time"
	"fmt"
)

type server struct {
	Sensor *sensor.Sensor
	pb.UnimplementedSensorServer
}


func (s *server) TempSensor(req *pb.SensorRequest, stream pb.Sensor_TempSensorServer) error {
	for {
		time.Sleep(time.Second * 5)
		temp := s.Sensor.GetTempSensor()
		err := stream.Send(&pb.SensorResponse{Value: temp})
		if err != nil{
			log.Println("Error sending metric message ", err)
		}
	}
	return nil;
}

func (s *server) HumiditySensor(req *pb.SensorRequest, stream pb.Sensor_HumiditySensorServer) error {
	for{
		time.Sleep(time.Second * 2)
		humd := s.Sensor.GetHumiditySensor()

		err := stream.Send(&pb.SensorResponse{Value:humd})
		if err != nil{
			log.Println("Error sending metric message ", err)
		}
	}
	return nil;
}



const (
	port = 8080
)
func main()  {
	sns := sensor.NewSensor()

	sns.StartMonitoring()

	addr := fmt.Sprintf("0.0.0.0:%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil{
		log.Fatalf("Error while listening : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSensorServer(s, &server{Sensor:sns})
	log.Printf("Starting server in port :%d\n", port)

	if err := s.Serve(lis); err != nil{
		log.Fatalf("Error while serving : %v", err)
	}
}