package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type HwAccepted struct {
	Id    int
	Grade int
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func currentDate() string {
	return time.Now().Format("2006-01-02")
}

func (a HwAccepted) logEvent(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s accepted %d %d\n", currentDate(), a.Id, a.Grade)
	return err
}

func (s HwSubmitted) logEvent(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s submitted %d \"%s\"\n", currentDate(), s.Id, s.Comment)
	return err
}

type OtusEvent interface {
	logEvent(io.Writer) error
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	if err := e.logEvent(w); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s Failed to log %v: %s\n", currentDate(), e, err)
	}
}
