package server

import (
	"context"
	"fmt"
	"log"
	"remy/pkg/ent/enttest"
	"remy/pkg/ent/event"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
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
	if err != nil {
		log.Println(err)
	} else {
		log.Println(type_date)
		for _, item := range type_date {
			log.Println(item)
		}
	}
}
