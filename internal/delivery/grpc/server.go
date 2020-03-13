package grpc

import (
	"context"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	config     *Config
}

// Config defines config for GRPC service
type Config struct {
	ServerOptions      []grpc.ServerOption
	Address            string
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

// New creates new grpc service
func NewServer(address string, opts ...Option) *Server {
	cfg := Config{Address: address}

	// default unary interceptors
	cfg.unaryInterceptors = append(cfg.unaryInterceptors,
		grpc_prometheus.UnaryServerInterceptor,
		grpc_validator.UnaryServerInterceptor(),
	)

	// default stream interceptors
	cfg.streamInterceptors = append(cfg.streamInterceptors,
		grpc_prometheus.StreamServerInterceptor,
	)

	for _, opt := range opts {
		cfg = opt(cfg)
	}

	cfg.ServerOptions = append(cfg.ServerOptions,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(cfg.streamInterceptors...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(cfg.unaryInterceptors...)),
	)

	grpc_prometheus.EnableHandlingTimeHistogram()
	s := grpc.NewServer(cfg.ServerOptions...)
	grpc_prometheus.Register(s)
	reflection.Register(s)

	return &Server{
		grpcServer: s,
		config:     &cfg,
	}
}

// Config return grpc service config
func (svc *Server) Config() *Config {
	return svc.config
}

// Address of grpc service
func (svc *Server) Address() string {
	return svc.config.Address
}

// Type of grpc service
func (svc *Server) Type() string {
	svcname := "grpc-service"
	return svcname
}

// Serve grpc server
func (svc *Server) Serve(ls net.Listener) error {
	return svc.Server().Serve(ls)
}

// Shutdown grpc server
func (svc *Server) Shutdown(ctx context.Context) error {
	svc.Server().GracefulStop()
	return nil
}

// Server return grpc server
func (svc *Server) Server() *grpc.Server {
	return svc.grpcServer
}

// Code returns service code to be used by admin page
func (svc Server) Code() string {
	return "grpc"
}

func (s *Server) Run() error {
	return s.runGRPC()
}

func (s *Server) runGRPC() error {
	listener, err := net.Listen("tcp", s.Address())
	if err != nil {
		return err
	}

	log.Printf("GRPC server running on port %v", s.Address())

	if err = s.grpcServer.Serve(listener); err != nil {
		return err
	}


	return nil
}
