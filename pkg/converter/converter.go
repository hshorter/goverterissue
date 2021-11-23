//go:generate go run github.com/jmattheis/goverter/cmd/goverter --packageName=converter --output=./pkg/converter/converter_impl.go github.com/hshorter/goverterissue/pkg/converter

package converter

import (
	"time"
)

// goverter:converter
// goverter:extend timeToTime
type DeviceConverter interface {
	ConvertToOutput(i Input) Output
}

type Input struct {
	Name string
	Time time.Time
}
type Output struct {
	Name string
	Time time.Time
}

func timeToTime(t time.Time) time.Time {
	return t
}
