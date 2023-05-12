package main

import (
	"fmt"
	proto "hellogrpc/proto"
	"io"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedUploaderServiceServer
}

func (s *Server) UploadFile(stream proto.UploaderService_UploadFileServer) error {
	for {
		fileData, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return errors.Wrapf(err, "failed unexpectadely while reading chunks from stream")
		}

		fmt.Println(fileData)
	}

	stream.SendMsg("done")
	return nil

}

func main() {
	server := &Server{}

	grpcServer := grpc.NewServer()
	proto.RegisterUploaderServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
