package main

import (
	"log"
	"context"
	"io"
	"google.golang.org/grpc"
	pb "client/charts"
	// "time"
	// "os"
	"sync"

    // "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/go-echarts/go-echarts/v2/opts"
    // "github.com/go-echarts/go-echarts/v2/types"
)

func HandleErr(err error)  {
	if err != nil{
		log.Fatalf("Some Error occurred %v", err)
	}
}

var wg sync.WaitGroup

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	HandleErr(err)
	defer conn.Close()

	client := pb.NewChartDataServiceClient(conn)
	
	stream, err := client.GetLineChartData(context.Background(), &pb.LineChartRequest{
		Title:"My Chart",
	})
	HandleErr(err)

	progressStream, err := client.GetProgressChartData(context.Background(), &pb.ProgressChartRequest {
	})
	HandleErr(err)

	// fmt.Println(reflect.TypeOf(client))
	sals := make([]opts.LineData,0)
	ages := make([]int64,0)
	progTitle := make([]string, 0)
	progValue := make([]int64, 0)

	for {
		wg.Add(2)
		go func() {
		res, err := stream.Recv()
		defer wg.Done()
		if err != nil {
			if err == io.EOF{
				return
			}
			log.Fatalf("failed to receive: %v", err)
		} else{
			log.Println(res.X, res.Y)
			sals = append(sals, opts.LineData{Value : res.Y})
			ages = append(ages, (res.Y) )
		}
		}()
		go func() {
			progRes, err :=  progressStream.Recv()
			defer wg.Done()
		if err != nil {
			if err == io.EOF{
				return
			}
			log.Fatalf("failed to receive: %v", err)
		} else{
			log.Printf("Title: %d, Value: %d", progRes.Title, progRes.Value)
			progTitle = append(progTitle, progRes.Title)
			progValue = append(progValue, progRes.Value)
		}
		}()
		wg.Wait()
		// time.Sleep(time.Second * 1)
		// line := charts.NewLine()

		// line.SetGlobalOptions(
		// 	charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		// 	charts.WithTitleOpts(opts.Title{Title: "My Chart",
		// 		Subtitle: "Test"}),
		// )

		// // log.Println(sals)

		// line.SetXAxis(ages).AddSeries("salaries", sals)

		// line.SetSeriesOptions(charts.WithLineChartOpts(
		// 	opts.LineChart{
		// 		Smooth: true,
		// 	}))

		// f, _ := os.Create("scatter.html")
		// line.Render(f)
	}
}