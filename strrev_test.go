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
		{in: "abåd", wantOut: "dåba"},
		// {in: "abo\u0301d", wantOut: "do\u0301ba"}, // multi-rune combining characters
		// {in: "ab👍🏽d", wantOut: "d👍🏽ba"}, // skin tone modifiers
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			out := Reverse(tc.in)
			if out != tc.wantOut {
				t.Errorf("Reverse %q got %q, want %q", tc.in, out, tc.wantOut)
			}

			outBytes := ReverseBytes([]byte(tc.in))
			if string(outBytes) != tc.wantOut {
				t.Errorf("ReverseBytes %q got %q, want %q", tc.in, string(outBytes), tc.wantOut)
			}
		})
	}
}
