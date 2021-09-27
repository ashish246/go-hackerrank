package main

import (
	"fmt"
	"math"
)

var totalUniversalSubarrays int
var uniSubarrList [][]int
func main() {

	// fmt.Println("LongestStablePriceCountArr => ", fetchLongestStablePrices([]int{3, 1,2,1,2,2,1,3,1,1,2,2,2,2}, 1))

	// fmt.Println("NoOfUniversalArray => ", countingUniversalSubarrays([]int{4,4,2,2,4,2}))
	processSubArrays([]int{4,4,2,2,4,2}, 0, 0)
	fmt.Println("totalUniversalSubarrays => ", totalUniversalSubarrays)
}

func processSubArrays(data []int, start, end int){

	if end == len(data) {
		return
	} else if start > end {
		processSubArrays(data, 0, end+1)
	} else {
		subArray := data[start:end+1]
		if len(subArray) % 2 == 0 {
			// fmt.Println("subArray => ", subArray)
			if countingUniversalSubarrays(subArray) {
				totalUniversalSubarrays++
				// FIXME: Remove the duplicate array from the count
				uniSubarrList = append(uniSubarrList, subArray)
				fmt.Println("Is universal subarray: TRUE")
			} else {
				fmt.Println("Is universal subarray: FALSE")
			}
		}
		processSubArrays(data, start+1, end)
	}
}

// countingUniversalSubarrays counts Universal Subarrays
func countingUniversalSubarrays(subArray []int) bool{
	fmt.Println("Checking for Unitversal subArray => ", subArray)
	prevValue  := 0
	prevValueCount  := 0

	newValue := 0
	newValueCount := 0
	for _, v := range subArray {
		if prevValue == 0 || (prevValue == v && newValue == 0) {
			if prevValue == 0 {
				prevValue = v
			}
			prevValueCount++
			prevValue = v
		} else {
			if newValue == 0 {
				newValue = v
			}
			// return false, if there is third value
			if newValue != v {
				return false
			}

			newValueCount++
		}
	}
	fmt.Println("prevValue: ", prevValue)
	fmt.Println("prevValueCount: ", prevValueCount)
	fmt.Println("newValue: ", newValue)
	fmt.Println("newValueCount: ", newValueCount)

	return prevValueCount == newValueCount
}

// fetchLongestStablePrices fetches longest stable market period
func fetchLongestStablePrices (data []int, x int) []int  {
	// data := [14]int{3, 1,2,1,2,2,1,3,1,1,2,2,2,2}
	// x := 1

	var stablePeriodArr [][]int 
	var prevPrice int
	var stablePeriod []int 
	lastIndex := len(data) - 1
	for i, currPrice := range data {
		if prevPrice == 0 {
			prevPrice = currPrice
		}
		// fmt.Println("prevPrice => ", prevPrice)
		// fmt.Println("currPrice => ", currPrice)
		diff := currPrice - prevPrice
		diffAbs := int(math.Abs(float64(diff)))
		// fmt.Println("diffAbs => ", diffAbs)
		// fmt.Println("diff abs => ", diffAbs)
		// fmt.Println("X => ", x)

		if diffAbs <= x {
			// fmt.Println("TRUE")
			stablePeriod = append(stablePeriod, currPrice)
		}
		if int(math.Abs(float64(diff))) > x {
			// fmt.Println("FALSE")
			stablePeriodArr = append(stablePeriodArr, stablePeriod)
			stablePeriod = nil
			stablePeriod = append(stablePeriod, currPrice)
		}
		prevPrice = currPrice
		// fmt.Println("stablePeriod => ", stablePeriod)
		if i ==  lastIndex {
			stablePeriodArr = append(stablePeriodArr, stablePeriod)
		}
	}

	// fmt.Println("stablePeriod => ", stablePeriod)
	// fmt.Println("stablePeriodArr => ", stablePeriodArr)
	longestStablePriceCount := 0
	var longestStablePriceCountArr []int
	// TODO: this loop can also be avoided by putting the logic in previous loop
	for _, v := range stablePeriodArr {
		if len(v) > longestStablePriceCount {
			longestStablePriceCount = len(v)
			longestStablePriceCountArr = v
		}
	}

	return longestStablePriceCountArr
}