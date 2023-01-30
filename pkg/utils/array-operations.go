package utils

// function which take an array of int, and return the product of all elements
func ArrayProduct(arr []int) int {
	var result int = 1
	for i := 0; i < len(arr); i++ {
		result *= arr[i]
	}
	return result
}

func ArrFill(arr []int, value int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = value
	}
	return arr
}
