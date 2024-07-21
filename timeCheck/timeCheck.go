package timecheck

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Timecheck() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	fmt.Println(n)
}

func generateRandomNum() int {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	return n
}

var stdin = bufio.NewReader(os.Stdin)

func inputIntValue() (int, error) {
	fmt.Printf("숫자값을 입력하세요")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

func Init() {
	randomNum := generateRandomNum()
	ctn := 1
	for {

		n, err := inputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n > randomNum {
				fmt.Println("입력한 숫자가 더 틉니다.", randomNum, ctn)
			} else if n < randomNum {
				fmt.Println("입력한 숫자가 더 작습니다.", randomNum, ctn)
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다, 시도한 횟수:", ctn)
				break
			}
			ctn++
		}
	}
}
