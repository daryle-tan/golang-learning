package array

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// _ is used to ignore the index
	for _, number := range numbers {
		sum += number
	}
	 return sum
 }


