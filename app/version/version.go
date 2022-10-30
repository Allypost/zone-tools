package version

import (
	"time"
)

const TimeFormat = "2006-01-02T15:04:05-0700"

var (
	buildTimestamp              = TimeFormat
	buildProgramName            = "inc-zone-soa"
	initialTime                 = time.Now()
	parsedTime       *time.Time = nil
)

func BuildTime() *time.Time {
	if parsedTime != nil {
		return parsedTime
	}

	if buildTimestamp == TimeFormat {
		parsedTime = &initialTime
		return parsedTime
	}

	t, err := time.Parse(TimeFormat, buildTimestamp)

	if err == nil {
		localTime := t.Local()
		parsedTime = &localTime
	} else {
		parsedTime = &initialTime
	}

	return parsedTime
}

func BuildProgramName() string {
	return buildProgramName
}

func BuildVersion() string {
	return BuildTime().Format(TimeFormat)
}
