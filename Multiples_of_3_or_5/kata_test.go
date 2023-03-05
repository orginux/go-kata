package kata

import "testing"

func TestMultiple3And5(t *testing.T) {
	multiple3And5Tests := []struct {
		number int
		want   int
	}{
		{
			number: 10,
			want:   23,
		},
		{
			number: 1000,
			want:   233168,
		},
		{
			number: 365,
			want:   30783,
		},
	}
	for _, tt := range multiple3And5Tests {
		got := Multiple3And5(tt.number)
		if got != tt.want {
			t.Errorf("\ngot %d\nwant: %d\n", got, tt.want)
		}
	}
}
