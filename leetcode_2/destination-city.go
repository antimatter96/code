package main

import "fmt"

func destCity(paths [][]string) string {
	dest := make(map[string]bool)
	start := make(map[string]bool)

	for _, v := range paths {
		start[v[0]] = true
		dest[v[1]] = true
	}

	for k := range start {
		delete(dest, k)
	}

	for k := range dest {
		return k
	}

	return ""
}

func driver__destCity() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
	fmt.Println(destCity([][]string{{"A", "Z"}}))
}

/*
func destCity(paths [][]string) string {
	mp := make(map[string]bool)

	for _, v := range paths {
		mp[v[0]] = true
	}

	for _, v := range paths {
		if _, ok := mp[v[1]]; !ok {
			return v[1]
		}
	}

	return ""
}
*/

/*
func destCity(paths [][]string) string {
	nextPoint := paths[0][1]
	lastPoint := nextPoint
	pathsRest := paths[1:]

	for {
		nextPoint = findRoot(pathsRest, nextPoint)
		if nextPoint == "" {
			break
		} else {
			lastPoint = nextPoint
		}
	}

	return lastPoint
}

func findRoot(paths [][]string, target string) string {
	for _, v := range paths {
		if target == v[0] {
			return v[1]
		}
	}

	return ""
}
*/
