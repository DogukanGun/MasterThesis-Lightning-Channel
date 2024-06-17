package tests

import (
	"MasterThesis/recorder"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Setenv("datasource", "root:@/test")
	db := recorder.Connect()
	recorder.Save(db, "message", map[string]interface{}{"Message": "Hi, this is a test"})
}
