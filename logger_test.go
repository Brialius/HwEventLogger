package logger

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestLogOtusEvent(t *testing.T) {
	type args struct {
		e OtusEvent
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			"Accepted 1",
			args{
				HwAccepted{
					5,
					9,
				},
			},
			fmt.Sprintf("%s accepted %d %d\n", currentDate(), 5, 9),
		},
		{
			"Accepted 2",
			args{
				HwAccepted{
					50,
					10,
				},
			},
			fmt.Sprintf("%s accepted %d %d\n", currentDate(), 50, 10),
		},
		{
			"Submitted 1",
			args{
				HwSubmitted{
					5,
					"",
					"please take a look at my homework",
				},
			},
			fmt.Sprintf("%s submitted %d \"%s\"\n", currentDate(), 5, "please take a look at my homework"),
		},
		{
			"Submitted 2",
			args{
				HwSubmitted{
					55,
					"",
					"take a look at my homework",
				},
			},
			fmt.Sprintf("%s submitted %d \"%s\"\n", currentDate(), 55, "take a look at my homework"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			LogOtusEvent(tt.args.e, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("LogOtusEvent() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
	t.Run("Stdout", func(t *testing.T) {
		out := os.Stdout
		LogOtusEvent(HwAccepted{1, 5}, out)
		LogOtusEvent(HwSubmitted{15, "code", "take a look at my homework"}, out)
	})
}
