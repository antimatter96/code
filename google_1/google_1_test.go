package main

import "testing"

func benchmarknoOfArrowsNew(arr []int, b *testing.B) {
	temp := make([]int, len(arr))
	for i := 0; i < b.N; i++ {
		copy(temp, arr)
		noOfArrowsNew(temp)
	}
}

func benchmarknoOfArrows(arr []int, b *testing.B) {
	temp := make([]int, len(arr))
	for i := 0; i < b.N; i++ {
		copy(temp, arr)
		noOfArrows(temp)
	}
}

func BenchmarkOld_1_Worst(b *testing.B) { benchmarknoOfArrows(testCaseBest1, b) }
func BenchmarkOld_1_Best(b *testing.B)  { benchmarknoOfArrows(testCaseWorst1, b) }
func BenchmarkNew_1_Worst(b *testing.B) { benchmarknoOfArrowsNew(testCaseBest1, b) }
func BenchmarkNew_1_Best(b *testing.B)  { benchmarknoOfArrowsNew(testCaseWorst1, b) }

func BenchmarkOld_1000_Worst(b *testing.B) { benchmarknoOfArrows(testCaseWorst1000, b) }
func BenchmarkOld_1000_Best(b *testing.B)  { benchmarknoOfArrows(testCaseBest1000, b) }
func BenchmarkNew_1000_Worst(b *testing.B) { benchmarknoOfArrowsNew(testCaseWorst1000, b) }
func BenchmarkNew_1000_Best(b *testing.B)  { benchmarknoOfArrowsNew(testCaseBest1000, b) }
