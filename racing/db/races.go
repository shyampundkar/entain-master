package db

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shyampundkar/entain-master/racing/proto/racing"
)

const (
	STATUS_OPEN   = "OPEN"
	STATUS_CLOSED = "CLOSED"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequest) ([]*racing.Race, error)

	// Get will return a race by id.
	Get(request *racing.GetRaceRequest) (*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

// Get a race
func (r *racesRepo) Get(request *racing.GetRaceRequest) (*racing.Race, error) {
	var clauses []string

	query := getRaceQueries()[racesList]

	clauses = append(clauses, " id="+strconv.FormatInt(request.Id, 10))

	query = applyWhereClause(clauses, query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	races, err := r.scanRaces(rows)
	if err != nil {
		return nil, err
	}

	if len(races) == 0 {
		return nil, errors.New("race not found")
	}

	return races[0], nil
}

// List the races
func (r *racesRepo) List(request *racing.ListRacesRequest) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args = r.applyFilter(query, request.Filter)

	query = applyOrderBy(query, request.Orderby)

	log.Printf("Query:%v \n", query)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

// Apply sorting & ordering clause to the query
func applyOrderBy(query string, orderBy string) string {
	if len(strings.TrimSpace(orderBy)) != 0 {
		query += " ORDER BY " + orderBy
	}
	return query
}

func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter) (string, []interface{}) {
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

	// Get race visibility filter condition
	raceFilter := getRaceVisibilityFilter(filter.OptionalRaceVisibility)
	if len(raceFilter) != 0 {
		// Keep condition to become part of Where clause later
		clauses = append(clauses, raceFilter)
	}

	query = applyWhereClause(clauses, query)

	return query, args
}

func applyWhereClause(clauses []string, query string) string {

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	return query
}

// Get race visibility filter criteria from the race_visibility
func getRaceVisibilityFilter(race_visibility racing.ListRacesRequestFilter_Visibility) string {
	switch race_visibility {
	case racing.ListRacesRequestFilter_HIDDEN:
		return "visible = false"
	case racing.ListRacesRequestFilter_VISIBLE:
		return "visible = true"
	case racing.ListRacesRequestFilter_SHOW_ALL:
		return ""
	default:
		log.Printf("invalid value for filter.RaceVisibility:%v, Type: %T\n", race_visibility, race_visibility)
	}
	return ""
}

func (m *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts

		race.Status = getRaceStatus(advertisedStart)

		races = append(races, &race)
	}

	return races, nil
}

// race.Status field that is derived based on their race.advertised_start_time`'s.
// The status is simply, `OPEN` or `CLOSED`.
// All races that have an `advertised_start_time` in the past will reflect `CLOSED`.
func getRaceStatus(advertisedStart time.Time) string {
	if advertisedStart.After(time.Now()) {
		return STATUS_OPEN
	} else {
		return STATUS_CLOSED
	}
}
