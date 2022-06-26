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
