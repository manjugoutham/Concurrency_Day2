package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var nums []float64

	fmt.Println("Enter a series of floating-point numbers (press enter to stop):")
	var num float64
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		nums = append(nums, num)
	}
	n := len(nums)
	m := n / 4
	subarrays := [4][]float64{
		nums[0:m],
		nums[m : 2*m],
		nums[2*m : 3*m],
		nums[3*m:],
	}

	var wg sync.WaitGroup
	for i, subarray := range subarrays {
		wg.Add(1)
		go func(i int, subarray []float64) {
			defer wg.Done()
			fmt.Printf("Sorting subarray %d: %v\n", i, subarray)
			sort.Float64s(subarray)
		}(i, subarray)
	}

	wg.Wait()
	sorted := make([]float64, n)
	for i := 0; i < n; i += m {
		copy(sorted[i:i+m], subarrays[i/m])
	}
	sort.Float64s(sorted)
	fmt.Println("Sorted list:", sorted)
}
