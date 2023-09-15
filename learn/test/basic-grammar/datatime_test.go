package golearning_test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	t.Logf("%s \n", now.Format("2006-01-02 15:04:05")) // 2023-09-15 14:13:59

	t1Str := "2023-03-14T08:52:32+00:00"
	t1, _ := time.Parse(time.RFC3339, t1Str)
	t.Logf("%#v format: %s\n", t1, t1.Local().Format("2006-01-02 15:04:05"))
	t.Logf("origin timezone format: %s\n", t1.Format("2006-01-02 15:04:05"))
	t.Logf("local  timezone format: %s\n", t1.Local().Format("2006-01-02 15:04:05"))
	t.Logf("8zone  timezone format: %s\n", t1.In(time.FixedZone("8zone", 8*60*60)).Format("2006-01-02 15:04:05"))
	t.Logf("0zone  timezone format: %s\n", t1.In(time.FixedZone("0zone", 0)).Format("2006-01-02 15:04:05"))
}

func TestDurition(t *testing.T) {
	hour := time.Hour.Seconds()
	t.Logf("1 hour seconds: %d", int(hour))
}
