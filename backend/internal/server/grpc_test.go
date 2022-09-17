package server

import (
	"context"
	"fmt"
	"log"
	"reflect"
	api "remy/api/pb"
	"remy/internal/transformer"
	"remy/pkg/ent"
	"remy/pkg/ent/enttest"
	"remy/pkg/ent/event"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// _driver = "sqlite3"
	// _url    = "file:ent?mode=memory&cache=shared&_fk=1"
	_driver = "mysql"
	_url    = "root:P@ssw0rd@tcp(localhost:3306)/remy?parseTime=True"
)

func TestServer_ListEvents(t *testing.T) {
	type records struct {
		Date *time.Time
	}

	var (
		ctx    = context.Background()
		client = enttest.Open(t, _driver, _url)

		err       error
		type_date []*records
	)
	query := client.Debug().Event.Query()
	var date_type = "DATE"

	err = query.Modify(func(s *sql.Selector) {
		s.Select(fmt.Sprintf("DATE(%s)", event.FieldStart)).GroupBy(fmt.Sprintf("%s(%s)", date_type, event.FieldStart))
	}).Scan(ctx, &type_date)
	var res []*api.ListEvents
	for _, item := range type_date {
		events := query.Where(event.StartEQ(*item.Date)).AllX(ctx)
		res = append(res, &api.ListEvents{
			Event: transformer.EventsToMessage(events),
			Time:  timestamppb.New(*item.Date),
		})
	}
	if err != nil {
		log.Println(err)
	} else {
		log.Println(type_date)
		for _, item := range type_date {
			log.Println(item)
		}
	}
}

func TestServer_CreateEvent(t *testing.T) {
	
}
