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
