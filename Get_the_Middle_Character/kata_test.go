package kata

import "testing"

func TestGetMiddle(t *testing.T) {
	getMiddleTests := []struct {
		s    string
		want string
	}{
		{
			s:    "test",
			want: "es",
		},
		{
			s:    "testing",
			want: "t",
		},
		{
			s:    "A",
			want: "A",
		},
		{
			s:    "",
			want: "",
		},
	}

	for _, tt := range getMiddleTests {
		got := GetMiddle(tt.s)
		if got != tt.want {
			t.Errorf("\ngot: %s\nwant: %s\n", got, tt.want)
		}
	}
}
