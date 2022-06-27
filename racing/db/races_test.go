package db

import (
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
