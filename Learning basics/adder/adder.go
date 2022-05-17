package integers

import "fmt"

func Add(num1, num2 int) int {
	return num1 + num2
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
