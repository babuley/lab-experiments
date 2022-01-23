package models

const (
	degreePerMinuteWithinOneHour float32 = 5 * 6.0 / 60.0
	degreePerSingleHour          float32 = 360.0 / 12.0
	degreePerMinute              float32 = 360.0 / 60.0
)

type Clock struct {
	Hour   int
	Minute int
}

var (
	times []*Clock
)

func (clock Clock) SolveAngle() (float32, error) {
	hourAngle := (float32(clock.Minute) * degreePerMinuteWithinOneHour) +
		(float32(clock.Hour) * degreePerSingleHour)
	minuteAngle := float32(clock.Minute) * degreePerMinute
	return hourAngle - minuteAngle, nil
}
