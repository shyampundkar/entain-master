package db

import (
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shyampundkar/entain-master/sports/proto/sports"
)

const (
	STATUS_OPEN   = "OPEN"
	STATUS_CLOSED = "CLOSED"
)

// EventsRepo provides repository access to Sport events.
type EventsRepo interface {
	// Init will initialise our events repository.
	Init() error

	// List will return a list of events.
	List(filter *sports.ListEventsRequest) ([]*sports.Event, error)
}

type eventsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewEventsRepo creates a new events repository.
func NewEventsRepo(db *sql.DB) EventsRepo {
	return &eventsRepo{db: db}
}

// Init prepares the event repository dummy data.
func (e *eventsRepo) Init() error {
	var err error

	e.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy events.
		err = e.seed()
	})
	return err
}

// List the events
func (e *eventsRepo) List(request *sports.ListEventsRequest) ([]*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportsQueries()[eventsList]

	query, args = e.applyFilter(query, request.Filter)

	query = applyOrderBy(query, request.Orderby)

	log.Printf("Query:%v \n", query)
	rows, err := e.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return e.scanEvents(rows)
}

// Apply sorting & ordering clause to the query
func applyOrderBy(query string, orderBy string) string {
	if len(strings.TrimSpace(orderBy)) != 0 {
		query += " ORDER BY " + orderBy
	}
	return query
}

func (e *eventsRepo) applyFilter(query string, filter *sports.ListEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	// Get event visibility filter condition
	eventFilter := getEventVisibilityFilter(filter.OptionalEventVisibility)
	if len(eventFilter) != 0 {
		// Keep condition to become part of Where clause later
		clauses = append(clauses, eventFilter)
	}

	query = applyWhereClause(clauses, query)

	return query, args
}

// append where clauses to query
func applyWhereClause(clauses []string, query string) string {

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	return query
}

// Get event visibility filter criteria from the event_visibility
func getEventVisibilityFilter(event_visibility sports.ListEventsRequestFilter_Visibility) string {
	switch event_visibility {
	case sports.ListEventsRequestFilter_HIDDEN:
		return "visible = false"
	case sports.ListEventsRequestFilter_VISIBLE:
		return "visible = true"
	case sports.ListEventsRequestFilter_SHOW_ALL:
		return ""
	default:
		log.Printf("invalid value for filter.getEventVisibilityFilter:%v, Type: %T\n", event_visibility, event_visibility)
	}
	return ""
}

func (m *eventsRepo) scanEvents(
	rows *sql.Rows,
) ([]*sports.Event, error) {
	var events []*sports.Event

	for rows.Next() {
		var event sports.Event
		var advertisedStart time.Time

		if err := rows.Scan(&event.Id, &event.MeetingId, &event.Name, &event.Number, &event.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		event.AdvertisedStartTime = ts

		event.Status = getEventStatus(advertisedStart)

		events = append(events, &event)
	}

	return events, nil
}

// event.Status field that is derived based on their event.advertised_start_time`'s.
// The status is simply, `OPEN` or `CLOSED`.
// All events that have an `advertised_start_time` in the past will reflect `CLOSED`.
func getEventStatus(advertisedStart time.Time) string {
	// TODO:Get the time zone from consumer
	// We are comparing with server UTC time.
	// Ideally we should compare advertisedStart with consumer provided UTC time.
	if advertisedStart.UTC().After(time.Now().UTC()) {
		return STATUS_OPEN
	} else {
		return STATUS_CLOSED
	}
}
