package main

import (
	"cli/pkg/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	"time"
)

func main() {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewThumbnailsClient(conn)

	linc := os.Args
	switch {
	case len(linc) < 2:
		fmt.Println("add YouTube URL")
	case len(linc) > 2:
		fmt.Println("Use key --async to download more then one url")
	}
	async := flag.Bool("async", false, "asynchronous downloading")

	flag.Parse()
	if *async {
		go func() {
			stream, _ := client.DownloadAsync(context.Background())
			if err := stream.Send(&wrapperspb.StringValue{Value: linc[1]}); err != nil {
				panic(err)
			}
		}()
	} else {
		image, err := client.Download(context.Background(), &wrapperspb.StringValue{Value: linc[1]})
		if err != nil {
			panic(err)
		}
		fmt.Println(image.String())
	}
}
