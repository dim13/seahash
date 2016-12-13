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
			d := New()
			d.Write([]byte(tc.s))
			if sum := d.Sum64(); sum != tc.n {
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
		t.Run(tc.b, func(t *testing.T) {
			a, b := New(), New()
			a.Write([]byte(tc.a))
			b.Write([]byte(tc.b))
			if a.Sum64() == b.Sum64() {
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
		t.Run(fmt.Sprintf("% x", tc.b), func(t *testing.T) {
			a, b := New(), New()
			a.Write(tc.a)
			b.Write(tc.b)
			if a.Sum64() == b.Sum64() {
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
