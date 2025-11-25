package lib

import (
	"fmt"
	"time"
)

func GetTimeNow(param string) string {
	location, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}

	currentTime := time.Now().In(location)

	switch param {
	case "timestime":
		fmt.Println("Local time:", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("Jakarta time:", currentTime.Format("2006-01-02 15:04:05"))
		return currentTime.Format("2006-01-02 15:04:05")
	case "date":
		return currentTime.Format("2006-01-02")
	case "date2":
		return currentTime.Format("020106")
	case "year":
		return fmt.Sprint(currentTime.Year())
	case "month":
		return fmt.Sprint(int(currentTime.Month()))
	case "month-name":
		return fmt.Sprint(currentTime.Month())
	case "day":
		return fmt.Sprint(currentTime.Day())
	case "hour":
		return string(currentTime.Hour())
	case "minutes":
		return string(currentTime.Minute())
	case "second":
		return string(currentTime.Second())
	case "unixmicro":
		return string(currentTime.UnixMicro())
	default:
		fmt.Println("masukan parameter")
		return ""
	}
}

func AddTime(year int, month int, days int) *string {
	currentTime := time.Now()
	time.LoadLocation("Asia/Jakarta")

	addtime := fmt.Sprint(currentTime.AddDate(year, month, days).Format("2006-01-02 15:04:05"))
	return &addtime
}

func FormatTime(timeString string) (string, error) {
	parsedTime, err := time.Parse("2002-10-02 15:04:05", timeString)
	if err != nil {
		return "", err
	}

	formattedTime := parsedTime.Format("October 2, 2022 at 15:04 PM")
	return formattedTime, nil
}

func FixEndDate(request string) (response string) {
	date, err := time.Parse("2006-01-02", request)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// Add one day to the date
	newDate := date.AddDate(0, 0, 1)

	// Format and print the new date
	newDateString := newDate.Format("2006-01-02")

	return newDateString
}

func FormatDatePtr(s *string) string {
	if s == nil {
		return ""
	}

	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return *s
	}

	return t.Format("2006-01-02 15:04:05")
}
