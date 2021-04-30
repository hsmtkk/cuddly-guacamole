package clock

import "fmt"

type Timer interface {
	Equal(t Timer) bool
	Add(minute int)
	Sub(minute int)
	Hour() int
	Minute() int
	String() string
}

type timerImpl struct {
	minute int
}

func New() Timer {
	return &timerImpl{minute: 0}
}

func (t *timerImpl) Equal(another Timer) bool {
	return t.Hour() == another.Hour() && t.Minute() == another.Minute()
}

func (t *timerImpl) Add(minute int) {
	t.minute += minute
	t.minute %= 60 * 24
}

func (t *timerImpl) Sub(minute int) {
	t.minute -= minute
	for {
		if t.minute < 0 {
			t.minute += 60 * 24
		} else {
			break
		}
	}
}

func (t *timerImpl) Hour() int {
	return t.minute / 60
}

func (t *timerImpl) Minute() int {
	return t.minute % 60
}

func (t *timerImpl) String() string {
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}
