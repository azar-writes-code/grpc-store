package main

import (
	"log"
	"context"
	"io"
	"google.golang.org/grpc"
	pb "client/charts"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Some Error occurred conn %v", err)
	}
	defer conn.Close()

	client := pb.NewChartDataServiceClient(conn)

	line, err := client.GetLineChartData(context.Background(), &pb.ChartRequest{})
	if err != nil {
		log.Fatalf("Some Error occurred line %v", err)
	}
	bar, err := client.GetBarChartData(context.Background(), &pb.ChartRequest{})
	if err != nil {
		log.Fatalf("Some Error occurred bar %v", err)
	}

	for {
		li, err := line.Recv()
		if err != nil {
			if err == io.EOF{
				return
			}
			log.Fatalf("failed to receive line: %v", err)
		} else{
			log.Println("======= Line Chart =========")
			log.Printf("X-axis Value : %v , Y-Axis Value : %v", li.X, li.Y)
		}

		ba, err := bar.Recv()
		if err != nil {
			if err == io.EOF{
				return
			}
			log.Fatalf("failed to receive bar: %v", err)
		} else{
			log.Println("======= Bar Chart =========")
			log.Printf("X-axis Value : %v , Y-Axis Value : %v", ba.X, ba.Y)
		}
	}
}