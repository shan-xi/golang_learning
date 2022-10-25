package morestrings

import "testing"

func TestReverseRUnes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"ABC", "CBA"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
