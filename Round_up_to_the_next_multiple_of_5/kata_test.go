package kata

import "testing"

func TestRoundToNext5(t *testing.T) {
	roundToNext5Tests := []struct {
		input int
		want  int
	}{
		{
			input: 0,
			want:  0,
		},
		{
			input: 2,
			want:  5,
		},
		{
			input: 12,
			want:  15,
		},
		{
			input: 30,
			want:  30,
		},
		{
			input: -5,
			want:  -5,
		},
	}
	for _, tt := range roundToNext5Tests {
		got := RoundToNext5(tt.input)
		if got != tt.want {
			t.Errorf("\ngot %d\nwant: %d\n", got, tt.want)
		}
	}
}
