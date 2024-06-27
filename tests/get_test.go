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
	var arr []Chat
	err := recorder.Get(db, "message", "", &arr)
	if err != nil {
		t.Fatalf("Failed to get data: %v", err)
	}
	assert.GreaterOrEqualf(t, len(arr), 1, "Must be greater or equal to 1")
}
