package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	api "sample_shecodes2022/api/pb"
	"sample_shecodes2022/internal/constants"

	_ "github.com/go-sql-driver/mysql"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type Service struct {
	logger *zap.Logger
	api.UnimplementedSampleServiceServer
}

func NewLogger(lv zapcore.Level, pretty bool) (*zap.Logger, error) {
	var c zap.Config
	var opts []zap.Option
	if pretty {
		c = zap.NewDevelopmentConfig()
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	} else {
		c = zap.NewProductionConfig()
	}

	level := zap.NewAtomicLevel()

	if err := level.UnmarshalText([]byte(lv.String())); err != nil {
		return nil, fmt.Errorf("could not parse log level %s", lv.String())
	}
	c.Level = level

	return c.Build(opts...)
}

func NewService() (*Service, error) {
	// Initial logger
	logger, errLog := NewLogger(zap.DebugLevel, true)
	if errLog != nil {
		return nil, errLog
	}

	return &Service{
		logger: logger,
	}, nil
}

func (s *Service) Close() {
	s.logger.Core().Sync()
}

func RunGRPC() {
	service, err := NewService()
	if err != nil {
		log.Fatalf("fail to create service: %s", err)
	}

	s := grpc.NewServer()
	api.RegisterSampleServiceServer(s, service)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *constants.Port))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	// Listen and serve port
	log.Printf("grpc server listening at %v", lis.Addr())

	go func() {
		log.Panicf("failed to serve: %v", s.Serve(lis))
		defer service.Close()
	}()
}
