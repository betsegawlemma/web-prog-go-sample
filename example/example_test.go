package example

import (
	"testing"
	"time"
)

func TestCustomDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2019, 12, 25, 10, 0, 0, 0, time.UTC),
			want: "25 Dec 2019 at 10:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "EAT",
			tm:   time.Date(2019, 12, 25, 10, 0, 0, 0, time.FixedZone("EAT", 3*60*60)),
			want: "25 Dec 2019 at 07:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := customDate(tt.tm)
			if got != tt.want {
				t.Errorf("want %q; got %q", tt.want, got)
			}
		})
	}
}
