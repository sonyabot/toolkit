package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var TestTools Tools
	s := TestTools.RandomString(10)
	if len(s) != 10 {
		t.Error("Wrong length random string returned")
	}
}
