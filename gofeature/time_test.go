package gostudy

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelForRangeSequence(t *testing.T) {
	ti, _ := time.ParseInLocation("2006-01-02", "2021-01-04", time.Local)
	t.Log(ti.Date())
	t.Log(ti.Clock())
	tt := ti.In(time.UTC)
	rs := tt.Format(time.RFC3339)
	t.Log(tt.Date())
	t.Log(tt.Clock())
	t.Log(rs)

}

func TestInLocation(t *testing.T) {
	ti, _ := time.ParseInLocation("2006-01-02", "2021-01-04", time.Local)
	t.Log(ti.Date())
	t.Log(ti.Clock())
	tt := ti.In(time.UTC)
	rs := tt.Format(time.RFC3339)
	t.Log(tt.Date())
	t.Log(tt.Clock())
	t.Log(rs)
}

func TestTruncate(t *testing.T) {
	today := time.Date(2022, 4, 24, 11, 55, 0, 0, time.Local)
	todayZeroReal := time.Date(2022, 4, 24, 0, 0, 0, 0, time.Local)
	todayZero := today.Truncate(24 * time.Hour)

	if todayZero.Equal(todayZeroReal) {
		t.Log("equal")
	} else {
		t.Error("not equal")
	}
}

func TestLocalTime(t *testing.T) {
	tm, _ := time.ParseInLocation("2006-01-02 15:04", "2020-10-19 11:34", time.Local)
	t.Log("unix: ", tm.Unix())
	t.Log("local unix: ", tm.Local().Unix())
	t.Log("utc unix: ", tm.UTC().Unix())
}

func TestUTC2Local(t *testing.T) {
	tm, _ := time.Parse("2006-01-02 15:04", "2020-10-19 11:34")
	t.Log(tm.Format(time.RFC3339))
	t.Log(tm.Local().Format(time.RFC3339))
	t.Log(tm.Unix())

	rfc, _ := time.Parse(time.RFC3339, "2020-10-19T19:34:00+08:00")
	t.Log(rfc.Format(time.RFC3339))
	t.Log(rfc.Local().Format(time.RFC3339))
	t.Log(rfc.Unix())

	time.Now().Local()
}

func TestTime(t *testing.T) {
	tm, _ := time.ParseInLocation("2006-01-02 15:04", "2020-10-19 11:34", time.Local)
	t.Log(tm.Unix())
}

func TestMonthDays(t *testing.T) {
	firstDay := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	days := firstDay.AddDate(0, 1, -1).Day()
	fmt.Println(days)
}

func TestNow(t *testing.T) {
	// t1, _ := time.ParseInLocation("2006-01-02 15:04", "2021-01-28 11:27", time.Local)
	t2, _ := time.Parse(time.RFC3339, "2021-01-28T11:27:00+00:00")
	fmt.Println(t2.Format(time.RFC3339))

	t3, _ := time.ParseInLocation(time.RFC3339, "2021-01-28T00:00:00+08:00", time.Local)
	fmt.Println(t3.Unix())
	fmt.Println(t3.Format(time.RFC3339))
	fmt.Println(t3.Local().Format(time.RFC3339))
}

func TestTruncateShouldOk(t *testing.T) {
	today := time.Date(2022, 4, 24, 11, 55, 0, 0, time.Local)
	todayZeroReal := time.Date(2022, 4, 24, 0, 0, 0, 0, time.Local)
	todayZero := TruncateDay(today)

	if todayZero.Equal(todayZeroReal) {
		t.Log("equal")
	} else {
		t.Error("not equal")
	}
}

func TruncateDay(t time.Time) time.Time {
	t = t.Truncate(24 * time.Hour)
	_, offset := t.Zone()

	return t.Add(-time.Duration(offset) * time.Second)
}

func TestTimeZone1(t *testing.T) {
	l, _ := time.LoadLocation("Africa/Niamey")
	today := time.Date(2022, 4, 24, 2, 0, 0, 0, l)
	fmt.Println(today.Zone())
	d := today.Truncate(2 * time.Hour)
	fmt.Println(d.Format("2006-01-02 15:04:05"))
}

func TestZY(t *testing.T) {
	l, _ := time.LoadLocation("Africa/Niamey")
	today := time.Date(2022, 4, 24, 1, 0, 0, 0, l)

	rs := today.In(l).Truncate(2 * time.Hour)
	fmt.Println(rs.Format("2006-01-02 15:04:05"))
}
