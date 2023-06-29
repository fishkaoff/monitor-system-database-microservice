package grpcserver

import (
	"context"
	"net"

	"github.com/fishkaoff/monitor-system-database-microservice/proto"
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

func (s *GRPCServer) Save(ctx context.Context, req *proto.SaveRequest) (*proto.SaveResponse, error) {
	storageResponse := s.svc.SaveSite(req.ChatID, req.Site)
	return &proto.SaveResponse{Message: storageResponse}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	storageResponse := s.svc.GetSites(req.ChatID)
	return &proto.GetResponse{
		Message: storageResponse,
	}, nil
}

func (s *GRPCServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	storageResponse := s.svc.DeleteSite(req.ChatID, req.Site)
	return &proto.DeleteResponse{
		Message: storageResponse,
	}, nil
}
