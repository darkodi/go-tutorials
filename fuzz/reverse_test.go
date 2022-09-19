package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
//     testcases := []struct {
//         in, want string
//     }{
//         {"Hello, world", "dlrow ,olleH"},
//         {" ", " "},
//         {"!12345", "54321!"},
//     }
//     for _, tc := range testcases {
//         rev := Reverse(tc.in)
//         if rev != tc.want {
//                 t.Errorf("Reverse: %q, want %q", rev, tc.want)
//         }
//     }
// }

// When fuzzing, you can’t predict the expected output, since you don’t have control over the inputs.
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		// Reversing a string twice preserves the original value
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		// The reversed string preserves its state as valid UTF-8.
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
