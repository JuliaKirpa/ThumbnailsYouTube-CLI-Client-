package main

import (
	"cli/pkg/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"os"
	"time"
)

func main() {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error grpc connection %s", err)
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
		//stream, _ := client.DownloadAsync(context.Background())
		//for _, value := range linc[1:] {
		//	if err := stream.Send(&wrapperspb.StringValue{Value: value}); err != nil {
		//		log.Fatalf("err to send: %s", err)
		//	}
		//}
		//channel := make(chan struct{})
		//go asncDownloadBdrectionalRPC(stream, channel)
		//if err := stream.CloseSend(); err != nil {
		//	log.Fatal(err)
		//}
		//<-channel
	} else {
		image, err := client.Download(context.Background(), &wrapperspb.StringValue{Value: linc[1]})
		if err != nil {
			log.Fatalf("error from download func %s", err)
		}
		fmt.Println(image.GetStatus(), image.GetId())
	}
}

//
//func asncDownloadBdrectionalRPC(stream proto.Thumbnails_DownloadAsyncClient, c chan struct{}) {
//	for {
//		images, err := stream.Recv()
//		if err == io.EOF {
//			break
//		}
//		fmt.Println(images.GetStatus(), images.GetId())
//	}
//	<-c
//}
