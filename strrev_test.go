package strrev

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		in      string
		wantOut string
	}{
		{in: "", wantOut: ""},
		{in: " ", wantOut: " "},
		{in: "a", wantOut: "a"},
		{in: "ab", wantOut: "ba"},
		{in: "abc", wantOut: "cba"},
		{in: "abcdefg", wantOut: "gfedcba"},
		{in: "ab丂d", wantOut: "d丂ba"},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			out := Reverse(tc.in)
			if out != tc.wantOut {
				t.Errorf("Reverse %q got %q, want %q", tc.in, out, tc.wantOut)
			}
		})
	}
}
