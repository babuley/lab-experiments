package models

import (
	"errors"
	"math"
)

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

func NewClock(hour int, minute int) (Clock, error) {
	c := Clock{
		Hour:   hour,
		Minute: minute,
	}
	if hour < 0 || hour > 24 || minute < 0 || minute > 60 {
		return c, errors.New("Illegal hour or minute error")
	}
	c.Angle, _ = c.SolveAngle()
	return c, nil
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
			clock, _ := NewClock(h, m)
			slice = append(slice, clock)
			m++
		}
		h++
	}
	return slice
}
