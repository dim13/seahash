package seahash

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	testCases := []struct {
		s string
		n uint64
	}{
		{"to be or not to be", 1988685042348123509},
		{"love is a wonderful terrible thing", 4784284276849692846},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			t.Parallel()
			if sum := SumString(tc.s); sum != tc.n {
				t.Errorf("got %v, want %v", sum, tc.n)
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	testCases := []struct {
		a, b string
	}{
		{"to be or not to be ", "to be or not to be"},
		{"jkjke", "jkjk"},
		{"ijkjke", "ijkjk"},
		{"iijkjke", "iijkjk"},
		{"iiijkjke", "iiijkjk"},
		{"iiiijkjke", "iiiijkjk"},
		{"iiiiijkjke", "iiiiijkjk"},
		{"iiiiiijkjke", "iiiiiijkjk"},
		{"iiiiiiijkjke", "iiiiiiijkjk"},
		{"iiiiiiiijkjke", "iiiiiiiijkjk"},
		{"ab", "bb"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.a, tc.b), func(t *testing.T) {
			t.Parallel()
			if SumString(tc.a) == SumString(tc.b) {
				t.Fail()
			}
		})
	}
}

func TestZeroSenitive(t *testing.T) {
	testCases := []struct {
		a, b []byte
	}{
		{[]byte{1, 2, 3, 4}, []byte{1, 0, 2, 3, 4}},
		{[]byte{1, 2, 3, 4}, []byte{1, 0, 0, 2, 3, 4}},
		{[]byte{1, 2, 3, 4}, []byte{1, 2, 3, 4, 0}},
		{[]byte{1, 2, 3, 4}, []byte{0, 1, 2, 3, 4}},
		{[]byte{0, 0, 0}, []byte{0, 0, 0, 0, 0}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%x %x", tc.a, tc.b), func(t *testing.T) {
			t.Parallel()
			if Sum(tc.a) == Sum(tc.b) {
				t.Fail()
			}
		})
	}
}

func BenchmarkHash(b *testing.B) {
	d := New()
	p := []byte("to be or not to be")
	for i := 0; i < b.N; i++ {
		d.Write(p)
	}
}
