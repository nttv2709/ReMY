package transformer

import (
	api "remy/api/pb"
	"remy/pkg/ent"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EventToMessage(event *ent.Event) (*api.Event) {
	return &api.Event{
		Id: event.ID,
		Title: event.Title,
		RangeTime: &api.RangeTime{
			Start: timestamppb.New(event.Start),
			End: timestamppb.New(event.End),
		},
	}
}

func EventsToMessage(events []*ent.Event) ([]*api.Event) {
	var res []*api.Event
	for _, item := range events {
		res = append(res, EventToMessage(item))
	}
	return res
}