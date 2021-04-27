package grain_test

import (
	"log"
	"math/big"
	"testing"

	"github.com/hsmtkk/ideal-waddle/grain"
	"github.com/stretchr/testify/assert"
)

func TestGrain(t *testing.T) {
	want := new(big.Int)
	want, ok := want.SetString("18446744073709551615", 10)
	if !ok {
		log.Fatal("fail")
	}
	got := grain.Grain()
	assert.Equal(t, want, got)
}

func TestTwoPower(t *testing.T) {
	assert.Equal(t, big.NewInt(1), grain.TwoPower(0))
	assert.Equal(t, big.NewInt(2), grain.TwoPower(1))
	assert.Equal(t, big.NewInt(4), grain.TwoPower(2))
	assert.Equal(t, big.NewInt(8), grain.TwoPower(3))
}
