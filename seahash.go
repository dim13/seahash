// Package seahash implements SeaHash algorithm
package seahash

import "hash"

// Size of SeaHash sum in bytes
const Size = 8

type digest struct {
	a, b, c, d uint64
	n          int
}

func (d *digest) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		var q []byte
		if len(p) > Size {
			p, q = p[:Size], p[Size:]
		}
		var x uint64
		for i := len(p) - 1; i >= 0; i-- {
			x <<= 8
			x |= uint64(p[i])
		}
		d.a, d.b, d.c, d.d = d.b, d.c, d.d, diffuse(d.a^x)
		p = q
	}
	d.n += n
	return n, nil
}

func (d *digest) Sum(b []byte) []byte {
	s := d.Sum64()
	for i := Size - 1; i >= 0; i-- {
		b = append(b, byte(s>>uint(8*i)))
	}
	return b
}

func (d *digest) Reset() {
	d.a = 0x16f11fe89b0d677c
	d.b = 0xb480a793d8e6c86c
	d.c = 0x6fe2e5aaf078ebc9
	d.d = 0x14f994a4c5259381
	d.n = 0
}

func (d *digest) Size() int {
	return Size
}

func (d *digest) BlockSize() int {
	return 1
}

func (d *digest) Sum64() uint64 {
	return diffuse(d.a ^ d.b ^ d.c ^ d.d ^ uint64(d.n))
}

func diffuse(x uint64) uint64 {
	x *= 0x6eed0e9da4d94a4f
	x ^= (x >> 32) >> (x >> 60)
	x *= 0x6eed0e9da4d94a4f
	return x
}

// New returns a new hash.Hash64 computing SeaHash
func New() hash.Hash64 {
	d := new(digest)
	d.Reset()
	return d
}

func Sum(p []byte) uint64 {
	d := New()
	d.Write(p)
	return d.Sum64()
}

func SumString(s string) uint64 {
	return Sum([]byte(s))
}
