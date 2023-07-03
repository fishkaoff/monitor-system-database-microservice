package grpcserver

import (
	"context"
	"net"

	"github.com/fishkaoff/monitor-system-proto-files/proto"
	"github.com/fishkaoff/monitor-system-database-microservice/service"
	"google.golang.org/grpc"
)

func GRPCServerRun(listenAddr string, svc service.Service) error {
	GRPCServer := NewGRPCServer(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterStorageServer(server, GRPCServer)
	return server.Serve(ln)
}


type GRPCServer struct {
	svc service.Service
	proto.UnimplementedStorageServer
}

func NewGRPCServer(svc service.Service) *GRPCServer {
	return &GRPCServer{svc: svc}
}

func (s *GRPCServer) SaveUrl(ctx context.Context, req *proto.SaveUrlRequest) (*proto.SaveUrlResponse, error) {
	storageResponse := s.svc.SaveSite(req.ChatID, req.Site)
	return &proto.SaveUrlResponse{Message: storageResponse}, nil
}

func (s *GRPCServer) GetUrl(ctx context.Context, req *proto.GetUrlRequest) (*proto.GetUrlResponse, error) {
	storageResponse := s.svc.GetSites(req.ChatID)
	return &proto.GetUrlResponse{
		Message: storageResponse,
	}, nil
}

func (s *GRPCServer) DeleteUrl(ctx context.Context, req *proto.DeleteUrlRequest) (*proto.DeleteUrlResponse, error) {
	storageResponse := s.svc.DeleteSite(req.ChatID, req.Site)
	return &proto.DeleteUrlResponse{
		Message: storageResponse,
	}, nil
}


func (s *GRPCServer) SaveUser(ctx context.Context, req *proto.SaveUserRequest) (*proto.SaveUserResponse, error) {
	storageResponse := s.svc.SaveUser(req.ChatID, req.Token)
	return &proto.SaveUserResponse{
		Message: storageResponse,
	}, nil
}