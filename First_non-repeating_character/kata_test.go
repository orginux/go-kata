package kata

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	firstNonRepeatingTests := []struct {
		str  string
		want string
	}{
		{
			str:  "stress",
			want: "t",
		},
		{
			str:  "ej9;HY9uvSTH:fd0QRltQ.E",
			want: "j",
		},
		{
			str:  "eRVn39eg15EqnS;K tOelobr4k9bclWYqY0yxOX",
			want: "V",
		},
		{
			str:  "sTreSS",
			want: "T",
		},
		{
			str:  "",
			want: "",
		},
	}
	for _, tt := range firstNonRepeatingTests {
		got := FirstNonRepeating(tt.str)
		if got != tt.want {
			t.Errorf("\ngot: %s\nwant: %s\n", got, tt.want)
		}
	}
}
