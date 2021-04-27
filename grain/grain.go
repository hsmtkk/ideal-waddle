package grain

import (
	"math/big"
)

func Grain() *big.Int {
	s := new(big.Int)
	for i := 0; i < 64; i++ {
		t := TwoPower(i)
		s.Add(s, t)
	}
	return s
}

func TwoPower(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	p := big.NewInt(1)
	for i := 0; i < n; i++ {
		p.Mul(p, big.NewInt(2))
	}
	return p
}
