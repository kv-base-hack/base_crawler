package mathutil

import "math/big"

// MinInt64 returns the minimum number between two given params.
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MinUint64 returns the minimum number between two given params.
func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func MaxBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) >= 0 {
		return a
	}
	return b
}

func MinBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) >= 0 {
		return b
	}
	return a
}

func AddBigInt(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}
