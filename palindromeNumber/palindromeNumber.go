package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	n := 0
	mapStore := make(map[int]int)

	for x != 0 {
		mapStore[n] = x % 10
		x /= 10
		n += 1
	}

	for i := 0; i < n/2; i++ {
		if mapStore[i] != mapStore[n-1-i] {
			return false
		}

	}

	return true

}

func main() {
	result := false
	result = isPalindrome(12321)
	fmt.Println(result)
}
