package main

import (
	"context"
	proto "hellogrpc/proto"
	"io"
	"os"

	"google.golang.org/grpc"
)

const chunkSize = 64 * 1024

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := proto.NewUploaderServiceClient(conn)

	// client.UploadFile()

	file, err := os.Open("/Users/max/Downloads/GoodAccess_3.43.dmg")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, chunkSize)
	uploader, err := client.UploadFile(context.TODO())
	if err != nil {
		panic(err)
	}

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		uploader.Send(&proto.UploadRequestType{
			Content:  buf[:n],
			Filename: "goodaccess.dmg",
		})
	}
}
