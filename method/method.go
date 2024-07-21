package method

import "fmt"

// type account struct {
// 	balance int
// }

// func withdrawFunc(a *account, amount int) {
// 	a.balance -= amount
// }

// func (a *account) withdrawMethod(amount int) {
// 	a.balance -= amount
// }

// func Exam() {
// 	a := &account{100}
// 	withdrawFunc(a, 30)
// 	a.withdrawMethod(30)
// 	fmt.Printf("%d \n", a.balance)
// }

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func Receiver() {
	var a myInt = 10

	fmt.Println(a.add(30))
	var b int = 20

	fmt.Println(myInt(b).add(50))
}

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func PointerAndValue() {
	var mainA *account = &account{100, "Jeo", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}
