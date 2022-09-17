package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	api "remy/api/pb"
	"remy/internal/constants"
	"remy/internal/database"
	"remy/internal/transformer"
	"remy/pkg/ent"
	"remy/pkg/ent/event"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		ent:    ent,
	}, nil
}

func (s *Server) Close() {
	s.logger.Core().Sync()
}

func (s *Server) CreateEvent(ctx context.Context, request *api.CreateEventRequest) (*api.CreateEventReply, error) {
	// return nil, nil
	log.Println("Request: ", request)
	event, err := s.ent.Event.Create().SetXPos(request.Location.X).SetYPos(request.Location.Y).SetStart(request.RangeTime.Start.AsTime()).SetEnd(request.RangeTime.End.AsTime()).SetTitle(request.Title).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &api.CreateEventReply{
		Id: event.ID,
	}, nil
}

func (s *Server) DeleteEvent(ctx context.Context, request *api.DeleteEventRequest) (*api.DeleteEventReply, error) {
	err := s.ent.Event.DeleteOneID(request.Id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &api.DeleteEventReply{}, nil
}

func (s *Server) UpdateEvent(ctx context.Context, request *api.UpdateEventRequest) (*api.UpdateEventReply, error) {
	_, err := s.ent.Event.Update().SetXPos(request.Event.Location.X).SetYPos(request.Event.Location.Y).SetStart(request.Event.RangeTime.Start.AsTime()).SetEnd(request.Event.RangeTime.End.AsTime()).SetTitle(request.Event.Title).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &api.UpdateEventReply{}, nil
}

func (s *Server) ListEvents(ctx context.Context, request *api.ListEventsRequest) (*api.ListEventsReply, error) {
	query := s.ent.Debug().Event.Query().Where()
	var date_type string
	switch *&request.DateType {
	case api.ListEventsRequest_DAY:
		date_type = "DATE"
	case api.ListEventsRequest_MONTH:
		date_type = "MONTH"
	case api.ListEventsRequest_YEAR:
		date_type = "YEAR"
	}
	type records struct {
		Date *time.Time
	}
	var type_date []*records
	err := query.Modify(func(s *sql.Selector) {
		s.Select(fmt.Sprintf("DATE(%s)", event.FieldStart)).GroupBy(fmt.Sprintf("%s(%s)", date_type, event.FieldStart))
	}).Scan(ctx, &type_date)
	if err != nil {
		return nil, err
	}
	var res []*api.ListEvents
	for _, item := range type_date {
		events := s.ent.Debug().Event.Query().Where(func(s *sql.Selector) {
        s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", event.FieldStart), item.Date))
    }).AllX(ctx)
		res = append(res, &api.ListEvents{
			Event: transformer.EventsToMessage(events),
			Time:  timestamppb.New(*item.Date),
		})
	}

	return &api.ListEventsReply{
		ListEvents: res,
	}, nil
}

func (s *Server) GetRemind(ctx context.Context, request *api.GetRemindRequest) (*api.GetRemindReply, error) {
	query, err := s.ent.Debug().Event.Query().Order(ent.Asc(event.FieldStart)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, q := range query {
		log.Println(q)
	}
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
