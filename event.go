package event

import (
	"errors"
	"unicode/utf8"

	"github.com/xustella/calender"
)

type Event struct {
	title string
	calender.Date
}

func (e *Event) SetTitle(t string) error {
	if utf8.RuneCountInString(t) > 25 {
		return errors.New("Title too long")
	}
	e.title = t
	return nil

}

func (e *Event) Title() string {
	return e.title
}
