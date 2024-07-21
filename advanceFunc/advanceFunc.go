package advancefunc

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func Sum() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
}

// func Print(args ...interface{}) string {
// 	for _, arg = range args {
// 		switch f := arg.(type) {
// 		case bool:
// 			val := arg.(bool)

// 		}
// 	}
// }

// func add(a, b int) int {
// 	return a + b
// }

// func mul(a, b int) int {
// 	return a * b
// }

type opFunc = func(int, int) int

func GetOperator(op string) opFunc {
	if op == "+" {
		// return add
		return func(a, b int) int {
			return a + b
		}
	}
	if op == "*" {
		// return mul
		return func(a, b int) int {
			return a * b
		}
	}
	return nil
}
