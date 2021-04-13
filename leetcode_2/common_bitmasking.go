package main

func bitMasker() {
	var noOfDigits int
	n := 1 << noOfDigits
	var z int

	for digit := 0; digit < noOfDigits; digit++ {
		z = 1 << digit
		for i := 0; i < n; i++ {
			if (i & z) > 0 {
				// j bit is set for number i
			}
		}
	}
}
