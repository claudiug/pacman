package notyet

import "fmt"

func main1() {
	fmt.Printf("Hello world!\n")
	value := 10

	result := changeValue(value)
	fmt.Println(result)
	fmt.Println(value)
	fmt.Println(sum(10, 4))
	fmt.Println(swap("claudiu", "garba"))
	fmt.Println(addUntil(10))
}

func changeValue(arg0 int) int {
	arg0 = 100
	return arg0
}

func sum(x int, y int) int {
	result := x + y
	return result
}

func swap(first, second string) (string, string) {
	return second, first
}

func addUntil(number int) int {
	result := 0
	for i := 0; i < number; i++ {
		result += i
	}
	return result
}
