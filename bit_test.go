package algorithm

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestHammingWeight(t *testing.T) {
	_ = hammingWeight(2)
}

func TestMain1(t *testing.T) {
	main1()
}

func TestReverseBits(t *testing.T) {
	_ = reverseBits(2)
}

func TestHammingDistance(t *testing.T) {
	assert.Equal(t, hammingDistance(1, 4), 2)
	assert.Equal(t, hammingDistance(3, 1), 1)
}

func TestFindComplement(t *testing.T) {
	assert.Equal(t, findComplement(5), 2)
	assert.Equal(t, findComplement(1), 0)
}

func TestFindComplement2(t *testing.T) {
	findComplement2(1)
	findComplement2(2)
}
