package tools

import "time"

func IsToday(t time.Time) bool {
	timeNow := time.Now()
	t_zero := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, t.Location())
	dayDuration, _ := time.ParseDuration("+24h")
	t_one := t_zero.Add(dayDuration)
	if t.Sub(t_zero) > 0 && t.Sub(t_one) < 0 {
		return true
	}
	return false
}

func IsYesterday(t time.Time) bool {
	timeNow := time.Now()
	t_zero := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, t.Location())
	dayAgo, _ := time.ParseDuration("-24h")
	t_minus_one := t_zero.Add(dayAgo)
	if t.Sub(t_minus_one) > 0 && t.Sub(t_zero) < 0 {
		return true
	}
	return false
}

func GetTodayStartTime() time.Time {
	timeNow := time.Now()
	t_zero := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)
	return t_zero
}
