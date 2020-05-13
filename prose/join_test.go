package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

const (
	apple              = "Apple"
	orange             = "Orange"
	pear               = "pear"
	appleAndOrange     = apple + " and " + orange
	appleOrangeAndPear = apple + ", " + orange + ", and " + pear
)

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{apple}, want: apple},
		testData{list: []string{apple, orange}, want: appleAndOrange},
		testData{list: []string{apple, orange, pear}, want: appleOrangeAndPear},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf(errorString(test.list, got, test.want))
		}
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("\nJoinWithCommas(%#v) \n got: \"%s\" \n want: \"%s\"", list, got, want)
}
