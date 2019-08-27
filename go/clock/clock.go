package clock

import "fmt"

// Clock ...
type Clock struct {
	Hour   int
	Minute int
}

// New clock
func New(hour, minute int) Clock {
	m := (hour*60 + minute) % (60 * 24)
	if m < 0 {
		m += (60 * 24)
	}

	c := Clock{
		Hour:   m / 60,
		Minute: m % 60,
	}

	return c
}

// String a stringer
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add clock minutes
func (c Clock) Add(minutes int) Clock {
	m := (c.Hour*60 + c.Minute + minutes) % (60 * 24)
	if m < 0 {
		m += (60 * 24)
	}

	c.Hour = m / 60
	c.Minute = m % 60
	return c
}

// Subtract clock minutes
func (c Clock) Subtract(minutes int) Clock {
	m := (c.Hour*60 + c.Minute - minutes) % (60 * 24)
	if m < 0 {
		m += (60 * 24)
	}

	c.Hour = m / 60
	c.Minute = m % 60
	return c
}
