package utils

import (
	"fmt"
	"strings"
	"time"
)

func CalcCrawledDateYear() string {
	return time.Now().Format("2006")
}

func CalcCrawledDateMonth() string {
	return time.Now().Format("2006-01")
}

func CalcCrawledDateWeek() string {
	year, week := time.Now().ISOWeek()

	return fmt.Sprintf("%d-0%d", year, week)
}

func CalcCrawledDateDay() string {
	return time.Now().Format("2006-01-02")
}

func CalcCrawledDateHour() string {
	return time.Now().Format("2006-01-02 15")
}

func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func DayPassByIndex(index int) string {
	return time.Now().AddDate(0, 0, -index).Format("2006-01-02")
}

// Crontab ...
type Crontab struct {
	Minute string
	Hour   string
	Day    string
	Month  string
	Week   string
}

// CalcCrawledDate ...
func CalcCrawledDate(schedule string) (result string) {
	params := strings.Split(schedule, " ")
	crontab := Crontab{
		params[0],
		params[1],
		params[2],
		params[3],
		params[4],
	}
	if crontab.Week != "*" {
		result = CalcCrawledDateWeek()
	} else if crontab.Hour == "*" && crontab.Day == "*" && crontab.Month == "*" && crontab.Week == "*" {
		result = CalcCrawledDateHour()
	} else if crontab.Day == "*" && crontab.Month == "*" && crontab.Week == "*" {
		result = CalcCrawledDateDay()
	} else if crontab.Month == "*" && crontab.Week == "*" {
		result = CalcCrawledDateMonth()
	} else if crontab.Week == "*" && crontab.Month != "*" {
		result = CalcCrawledDateYear()
	}

	return
}

// CalcCrawledCycle ...
func CalcCrawledCycle(schedule string) (result string) {
	params := strings.Split(schedule, " ")
	crontab := Crontab{
		params[0],
		params[1],
		params[2],
		params[3],
		params[4],
	}
	if crontab.Week != "*" {
		result = "每周"
	} else if crontab.Hour == "*" && crontab.Day == "*" && crontab.Month == "*" && crontab.Week == "*" {
		result = "每小时"
	} else if crontab.Day == "*" && crontab.Month == "*" && crontab.Week == "*" {
		result = "每天"
	} else if crontab.Month == "*" && crontab.Week == "*" {
		result = "每月"
	} else if crontab.Week == "*" && crontab.Month != "*" {
		result = "每年"
	}

	return
}
