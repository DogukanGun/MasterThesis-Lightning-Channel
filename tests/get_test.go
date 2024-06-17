package tests

import (
	"MasterThesis/recorder"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Chat struct {
	Message string
	Sender  string
}

func TestGet(t *testing.T) {
	t.Setenv("datasource", "root:@/test")
	db := recorder.Connect()
	arr := recorder.Get(db, "message", "", Chat{})
	assert.GreaterOrEqualf(t, len(arr), 1, "Must be greater or equal to 1")
}
