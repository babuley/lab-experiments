package models

import "math"

const (
	degreePerMinuteWithinOneHour float32 = 5 * 6.0 / 60.0
	degreePerSingleHour          float32 = 360.0 / 12.0
	degreePerMinute              float32 = 360.0 / 60.0
)

type Clock struct {
	Hour   int
	Minute int
	Angle  float32
}

func NewClock(hour int, minute int) Clock {
	if hour < 0 || hour > 24 || minute < 0 || minute > 60 {
		panic("Invalid hour or minute")
	}

	c := Clock{
		Hour:   hour,
		Minute: minute,
	}
	c.Angle, _ = c.SolveAngle()
	return c
}

func (clock Clock) SolveAngle() (float32, error) {
	hourAngle := (float32(clock.Minute) * degreePerMinuteWithinOneHour) +
		(float32(clock.Hour) * degreePerSingleHour)
	minuteAngle := float32(clock.Minute) * degreePerMinute
	return float32(math.Abs(float64(hourAngle - minuteAngle))), nil
}

func GetAllClocks() []Clock {
	slice := []Clock{}
	h := 0
	for h < 12 {
		m := 0
		for m < 60 {
			slice = append(slice, NewClock(h, m))
			m++
		}
		h++
	}
	return slice
}
