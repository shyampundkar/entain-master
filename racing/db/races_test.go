package db

import (
	"time"

	"github.com/shyampundkar/entain-master/racing/proto/racing"

	"github.com/stretchr/testify/assert"

	//nolint:golint
	"testing"
)

func Test_Race_Visibility_Hidden(t *testing.T) {
	visibilityFilter := getRaceVisibilityFilter(racing.ListRacesRequestFilter_HIDDEN)
	visibilityClause := "visible = false"
	assert.Equal(t, visibilityFilter, visibilityClause)
}

func Test_Race_Visibility_Visible(t *testing.T) {
	visibilityFilter := getRaceVisibilityFilter(racing.ListRacesRequestFilter_VISIBLE)
	visibilityClause := "visible = true"
	assert.Equal(t, visibilityFilter, visibilityClause)
}

func Test_Race_Visibility(t *testing.T) {
	visibilityFilter := getRaceVisibilityFilter(racing.ListRacesRequestFilter_SHOW_ALL)
	visibilityClause := ""
	assert.Equal(t, visibilityFilter, visibilityClause)
}

func Test_Race_ApplySortOrder(t *testing.T) {
	query := "select * from race"
	actual := applyOrderBy(query, "advertised_start_time desc")
	expected := query + " ORDER BY advertised_start_time desc"
	assert.Equal(t, expected, actual)
}

func Test_Race_Status_Open(t *testing.T) {
	pasttimestr := "2023-03-02T21:28:00Z"
	tm, _ := time.Parse(time.RFC3339, pasttimestr)
	raceStatus := getRaceStatus(tm.UTC())
	assert.Equal(t, STATUS_OPEN, raceStatus)
}

func Test_Race_Status_Closed(t *testing.T) {
	raceStatus := getRaceStatus(time.Now().UTC())
	assert.Equal(t, STATUS_CLOSED, raceStatus)
}
