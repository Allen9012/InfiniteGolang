package main

import (
	"testing"
	"time"
)

func TestCheckEditParam(t *testing.T) {
	var dtime time.Time
	var DtErr error
	if dtime, DtErr = time.ParseInLocation("2006-01-02 15:04:05", "", time.Local); DtErr != nil {
		t.Error(DtErr)
	}
	t.Logf("dtime %d", dtime.Unix())
}
