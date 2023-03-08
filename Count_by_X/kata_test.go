package kata

import (
	"reflect"
	"testing"
)

func TestCountBy(t *testing.T) {
	countByTests := []struct {
		x    int
		n    int
		want []int
	}{
		{
			x:    1,
			n:    10,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			x:    2,
			n:    5,
			want: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range countByTests {
		got := CountBy(tt.x, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("\ngot: %#v\nwant: %#v", got, tt.want)
		}
	}
}
