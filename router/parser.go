package router

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidDueDate = errors.New("invalid due date")

func parseDueDate(dueDate string) (time.Time, error) {
	dueDateSlice := strings.Split(dueDate, " ")
	if len(dueDateSlice) < 1 || len(dueDateSlice) > 2 {
		return time.Time{}, ErrInvalidDueDate
	}

	if t, err := time.Parse(time.RFC3339, dueDate); err == nil {
		return t, nil
	}

	date, err := parseDate(dueDateSlice[0])
	if err != nil {
		return time.Time{}, err
	}

	if len(dueDateSlice) == 2 {
		h, m, err := parseTime(dueDateSlice[1])
		if err != nil {
			return time.Time{}, err
		}
		date = date.Add(time.Hour*time.Duration(h) + time.Minute*time.Duration(m))
		if date.Before(time.Now()) {
			date = date.AddDate(1, 0, 0)
		}
		return date, nil
	} else {
		date = date.Add(time.Hour*23 + time.Minute*59)
		if date.Before(time.Now()) {
			date = date.AddDate(1, 0, 0)
		}
		return date, nil
	}
}

func parseDate(dateStr string) (date time.Time, err error) {
	if date, err = time.Parse("2006-1-2", dateStr); err == nil {
		return date, nil
	} else if date, err = time.Parse("2006/1/2", dateStr); err == nil {
		return date, nil
	} else if d, err := time.Parse("1/2", dateStr); err == nil {
		return time.Date(time.Now().Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local), nil
	} else if dateStr == "今日" || dateStr == "today" || dateStr == "tdy" {
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local), nil
	} else if dateStr == "明日" || dateStr == "tomorrow" || dateStr == "tmr" {
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local), nil
	} else if res := regexp.MustCompile(`^(\d+)日後$`).FindStringSubmatch(dateStr); len(res) == 2 && res[1] != "" {
		n, err := strconv.Atoi(res[1])
		if err != nil {
			return time.Time{}, err
		}
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+n, 0, 0, 0, 0, time.Local), nil
	} else if res := regexp.MustCompile(`^\+(\d+)$`).FindStringSubmatch(dateStr); len(res) == 2 && res[1] != "" {
		n, err := strconv.Atoi(res[1])
		if err != nil {
			return time.Time{}, err
		}
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+n, 0, 0, 0, 0, time.Local), nil
	} else if dateStr == "月" || dateStr == "火" || dateStr == "水" || dateStr == "木" || dateStr == "金" || dateStr == "土" || dateStr == "日" {
		temp := map[string]int{"月": 1, "火": 2, "水": 3, "木": 4, "金": 5, "土": 6, "日": 7}
		d := (temp[dateStr] - int(time.Now().Weekday()) + 7) % 7
		if d == 0 {
			d = 7
		}
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+d, 0, 0, 0, 0, time.Local), nil
	} else if dateStr == "月曜" || dateStr == "火曜" || dateStr == "水曜" || dateStr == "木曜" || dateStr == "金曜" || dateStr == "土曜" || dateStr == "日曜" {
		temp := map[string]int{"月曜": 1, "火曜": 2, "水曜": 3, "木曜": 4, "金曜": 5, "土曜": 6, "日曜": 7}
		d := (temp[dateStr] - int(time.Now().Weekday()) + 7) % 7
		if d == 0 {
			d = 7
		}
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+d, 0, 0, 0, 0, time.Local), nil
	} else if dateStr == "月曜日" || dateStr == "火曜日" || dateStr == "水曜日" || dateStr == "木曜日" || dateStr == "金曜日" || dateStr == "土曜日" || dateStr == "日曜日" {
		temp := map[string]int{"月曜日": 1, "火曜日": 2, "水曜日": 3, "木曜日": 4, "金曜日": 5, "土曜日": 6, "日曜日": 7}
		d := (temp[dateStr] - int(time.Now().Weekday()) + 7) % 7
		if d == 0 {
			d = 7
		}
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+d, 0, 0, 0, 0, time.Local), nil
	}
	return time.Time{}, ErrInvalidDueDate
}

func parseTime(timeStr string) (hour, minute int, err error) {
	if hm, err := time.Parse("15:04", timeStr); err == nil {
		return hm.Hour(), hm.Minute(), nil
	} else {
		return 0, 0, ErrInvalidDueDate
	}
}
