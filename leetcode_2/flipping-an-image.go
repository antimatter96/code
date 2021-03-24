package main

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)

	k := n / 2
	var x, y int

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			x = image[i][j]
			y = image[i][n-j-1]

			if x == 0 {
				image[i][n-j-1] = 1
			} else {
				image[i][n-j-1] = 0
			}
			if y == 0 {
				image[i][j] = 1
			} else {
				image[i][j] = 0
			}
		}

		if n%2 == 1 {
			if image[i][k] == 0 {
				image[i][k] = 1
			} else {
				image[i][k] = 0
			}
		}

	}

	return image
}
