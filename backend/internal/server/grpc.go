package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	api "remy/api/pb"
	"remy/internal/constants"
	"remy/internal/database"
	"remy/pkg/ent"

	_ "github.com/go-sql-driver/mysql"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type Server struct {
	logger *zap.Logger
	api.UnimplementedCalendarServiceServer
	ent *ent.Client
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

func NewServer(ent *ent.Client) (*Server, error) {
	// Initial logger
	logger, errLog := NewLogger(zap.DebugLevel, true)
	if errLog != nil {
		return nil, errLog
	}

	return &Server{
		logger: logger,
	}, nil
}

func (s *Server) Close() {
	s.logger.Core().Sync()
}

func (s *Server) CreateEvent(ctx context.Context, request *api.CreateEventRequest) (*api.CreateEventReply, error) {
	return nil, nil
}

func RunGRPC() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *constants.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	db, err := database.Init()
	if err != nil {
		log.Println(err)
	}

	service, err := NewServer(db)
	if err != nil {
		log.Fatalf("fail to create service: %s", err)
	}

	s := grpc.NewServer()
	api.RegisterCalendarServiceServer(s, service)
	log.Printf("grpc server listening at %v", lis.Addr())
	go func() {
		log.Fatalf("failed to serve: %v", s.Serve(lis))
	}()

	// GRPC cannot connect directly to web. Connect through grpcWeb
	// grpcWebServer := grpcweb.WrapServer(
	// 	s,
	// 	// Enable CORS
	// 	grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	// )
	// handler := func(res http.ResponseWriter, req *http.Request) {
	// 	grpcWebServer.ServeHTTP(res, req)
	// }

	// srv := &http.Server{
	// 	Handler: http.HandlerFunc(handler),
	// 	Addr:    fmt.Sprintf("0.0.0.0:%d", *constants.Port+1),
	// }

	// log.Printf("http server listening at %v", srv.Addr)
	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
