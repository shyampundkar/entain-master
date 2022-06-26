package db

import (
	"github.com/shyampundkar/entain-master/racing/proto/racing"

	"github.com/stretchr/testify/assert"

	//nolint:golint
	"testing"
)

func Test_Race_Visibility(t *testing.T) {
	visibilityFilter := getRaceVisibilityFilter(racing.ListRacesRequestFilter_hidden)
	visibilityClause := "visible = false"
	assert.Equal(t, visibilityFilter, visibilityClause)
}

func Test_Race_ApplySortOrder(t *testing.T) {
	query := "select * from race"
	actual := applyOrderBy(query, "advertised_start_time desc")
	expected := query + " ORDER BY advertised_start_time desc"
	assert.Equal(t, expected, actual)
}
