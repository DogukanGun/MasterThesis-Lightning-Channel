package tests

import (
	"MasterThesis/channel"
	"testing"
)

func TestSaveIpfs(t *testing.T) {
	t.Setenv("LIGHTHOUSE_KEY", "")
	channel.SaveToIpfs("Test.pdf")
}
