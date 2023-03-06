package kata

import "testing"

func TestBandNameGenerator(t *testing.T) {
	bandNameGeneratorTests := []struct {
		word string
		want string
	}{
		{
			word: "dolphin",
			want: "The Dolphin",
		},
		{
			word: "alaska",
			want: "Alaskalaska",
		},
		{
			word: "knife",
			want: "The Knife",
		},
		{
			word: "tart",
			want: "Tartart",
		},
	}
	for _, tt := range bandNameGeneratorTests {
		got := bandNameGenerator(tt.word)
		if got != tt.want {
			t.Errorf("\ngot %s\nwant: %s\n", got, tt.want)
		}

	}
}
