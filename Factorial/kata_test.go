package kata

import "testing"

func TestFactorial(t *testing.T) {
	factorialTests := []struct {
		n    int
		want int
	}{
		{
			n:    0,
			want: 1,
		},
		{
			n:    1,
			want: 1,
		},
		{
			n:    4,
			want: 24,
		},
		{
			n:    5,
			want: 120,
		},
	}
	for _, tt := range factorialTests {
		got := Factorial(tt.n)
		if got != tt.want {
			t.Errorf("\ngot %d\nwant: %d\n", got, tt.want)
		}
	}
}
