// Package seahash implements SeaHash algorithm
package seahash

import "hash"

// Size of SeaHash sum in bytes
const Size = 8

type digest struct {
	sum, a, b, c, d uint64
}

// New returns a new hash.Hash64 computing SeaHash
func New() hash.Hash64 {
	return &digest{
		0,
		0x16f11fe89b0d677c,
		0xb480a793d8e6c86c,
		0x6fe2e5aaf078ebc9,
		0x14f994a4c5259381,
	}
}

func diffuse(x uint64) uint64 {
	x *= 0x6eed0e9da4d94a4f
	x ^= (x >> 32) >> (x >> 60)
	x *= 0x6eed0e9da4d94a4f
	return x
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
	d.sum = diffuse(d.a ^ d.b ^ d.c ^ d.d ^ uint64(n))
	return n, nil
}

func (d *digest) Sum(b []byte) []byte {
	s := d.Sum64()
	for i := Size - 1; i >= 0; i-- {
		b = append(b, byte(s>>uint(8*i)))
	}
	return b
}

func (d *digest) Reset()         { d.sum = 0 }
func (d *digest) Size() int      { return Size }
func (d *digest) BlockSize() int { return 1 }
func (d *digest) Sum64() uint64  { return d.sum }
