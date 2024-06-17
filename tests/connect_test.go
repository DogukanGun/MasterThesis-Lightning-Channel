package tests

import (
	"MasterThesis/recorder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	t.Setenv("datasource", "root:@/test")
	db := recorder.Connect()
	err := db.Ping()
	assert.Equal(t, err == nil, true)
}
