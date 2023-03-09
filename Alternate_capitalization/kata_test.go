package kata

import (
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {

	capitalizeTests := []struct {
		st   string
		want []string
	}{
		{
			st:   "abcdef",
			want: []string{"AbCdEf", "aBcDeF"},
		},
		{
			st:   "codingisafunactivity",
			want: []string{"CoDiNgIsAfUnAcTiViTy", "cOdInGiSaFuNaCtIvItY"},
		},
	}
	for _, tt := range capitalizeTests {
		got := Capitalize(tt.st)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("\ngot: %#v\nwant: %#v", got, tt.want)
		}
	}
}
