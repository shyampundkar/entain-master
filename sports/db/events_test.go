package db

import (
	"time"

	"github.com/shyampundkar/entain-master/sports/proto/sports"

	"github.com/stretchr/testify/assert"

	//nolint:golint
	"testing"
)

func Test_Event_Visibility_Hidden(t *testing.T) {
	visibilityFilter := getEventVisibilityFilter(sports.ListEventsRequestFilter_HIDDEN)
	visibilityClause := "visible = false"
	assert.Equal(t, visibilityFilter, visibilityClause)
}

func Test_Event_Visibility_Visible(t *testing.T) {
	visibilityFilter := getEventVisibilityFilter(sports.ListEventsRequestFilter_VISIBLE)
	visibilityClause := "visible = true"
	assert.Equal(t, visibilityFilter, visibilityClause)
}

func Test_Event_Visibility(t *testing.T) {
	visibilityFilter := getEventVisibilityFilter(sports.ListEventsRequestFilter_SHOW_ALL)
	visibilityClause := ""
	assert.Equal(t, visibilityFilter, visibilityClause)
}

func Test_Event_ApplySortOrder(t *testing.T) {
	query := "select * from events"
	actual := applyOrderBy(query, "advertised_start_time desc")
	expected := query + " ORDER BY advertised_start_time desc"
	assert.Equal(t, expected, actual)
}

func Test_Event_Status_Open(t *testing.T) {
	pasttimestr := "2023-03-02T21:28:00Z"
	tm, _ := time.Parse(time.RFC3339, pasttimestr)
	eventStatus := getEventStatus(tm.UTC())
	assert.Equal(t, STATUS_OPEN, eventStatus)
}

func Test_Event_Status_Closed(t *testing.T) {
	eventStatus := getEventStatus(time.Now().UTC())
	assert.Equal(t, STATUS_CLOSED, eventStatus)
}

func Test_Event_Where_Clause(t *testing.T) {
	query := "select * from events"
	whereClause := " WHERE id = 1"
	var clauses []string
	clauses = append(clauses, "id = 1")
	applyWhereClause(clauses, query)
	assert.Equal(t, query+whereClause, applyWhereClause(clauses, query))
}
