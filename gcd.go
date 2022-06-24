package timer

func BinaryGcd(a, b int64) int64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	k := int64(0)
	for (a|b)&1 == 0 {
		a >>= 1
		b >>= 1
		k++
	}

	for (a & 1) == 0 {
		a >>= 1
	}

	for {
		for (b & 1) == 0 {
			b >>= 1
		}
		if a > b {
			a, b = b, a
		}
		b -= a
		b >>= 1
		if b == 0 {
			return a << k
		}
	}
}
