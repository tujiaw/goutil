package goutil

import (
	"fmt"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

func NowFormat() string {
	return time.Now().Format(TIME_FORMAT)
}

func FromNow(t time.Time) string {
	return FromNowUnix(t.Unix())
}

func FromNowUnix(t int64) string {
	minutes := (time.Now().Unix() - t) / 60
	if minutes <= 0 {
		return "刚刚"
	} else if minutes < 60 {
		return fmt.Sprintf("%d分钟前", minutes)
	}
	hours := minutes / 60
	if hours < 24 {
		return fmt.Sprintf("%d小时前", hours)
	}
	days := hours / 24
	if days < 30 {
		return fmt.Sprintf("%d天前", days)
	}
	months := days / 30
	if months < 12 {
		return fmt.Sprintf("%d个月前", months)
	}
	years := months / 12
	return fmt.Sprintf("%d年前", years)
}
