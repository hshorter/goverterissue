package mapping

import (
	"time"
)

// goverter:converter
// goverter:extend timeToTime
type DeviceConverter interface {
	ConvertToOutput(i Input) Output
}

type Input struct {
	name string
	time time.Time
}
type Output struct {
	name string
	time time.Time
}

func timeToTime(t time.Time) time.Time {
	return t
}
