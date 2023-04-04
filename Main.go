package main

import (
	"fmt"
	"sort"
)

func detect(m int, k int, arr []int) []int {
	var sumMin int
	var sumMax int
	var c int
	det := make([]int, k)
	for i := 0; i < k-1; i++ {
		sumMin = sumMin + arr[i]
		sumMax = sumMax + arr[k-1-i]
		fmt.Println(sumMax)
		if m <= sumMax && m >= sumMin {
			det[c] = i + 1
			c = c + 1
		}
	}
	return det
}

func findCombinations(n int, m int, k int, usedNums []int, currentCombination []int, allCombinations *[][]int) {
	if n == 0 && m == 0 {
		sort.Slice(currentCombination, func(i, j int) bool {
			return i > j
		})
		*allCombinations = append(*allCombinations, currentCombination)
		return
	}

	for i := 1; i <= k; i++ {
		if !contains(usedNums, i) && i <= m {
			newUsedNums := append(usedNums, i)
			newCombination := append(currentCombination, i)
			findCombinations(n-1, m-i, i-1, newUsedNums, newCombination, allCombinations)
		}
	}
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	var m int
	fmt.Println("Введите m ")
	fmt.Scan(&m)
	var k int
	fmt.Println("Введите k ")
	fmt.Scan(&k)
	det := make([]int, 0)
	var a int
	a = 1
	slice := make([]int, k)
	for i := 0; i < k; i++ {
		slice[i] = a
		a = a + 1
	}
	det = detect(m, k, slice)
	var Nu int
	for i := 0; i <= len(det); i++ {
		if det[i] == 0 {
			Nu = i
			break
		} else {
			Nu = len(det)
		}
	}
	fmt.Println(det)
	usedNums := make([]int, 0)
	currentCombination := make([]int, 0)
	allCombinations := make([][]int, 0)
	for i := 0; i < Nu; i++ {
		findCombinations(det[i], m, k, usedNums, currentCombination, &allCombinations)
	}
	fmt.Println(allCombinations)
}
