package main

import (
	"context"
	"fmt"
	pb "client/sensorpb"
	"log"
	"google.golang.org/grpc"
)

func main()  {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	t:= pb.NewSensorClient(conn)

	humidityStream, err := t.HumiditySensor(context.Background(), &pb.SensorRequest{})

	if err != nil {
		log.Fatal(err)
	}

	tempStream, err := t.TempSensor(context.Background(), &pb.SensorRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		td, err := tempStream.Recv()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Temp : ", td.Value)
		}
		hd, err := humidityStream.Recv()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Humidity : ", hd.Value)
		}
		fmt.Println()
	}
}