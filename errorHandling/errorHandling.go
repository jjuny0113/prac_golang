package errorhandling

import (
	"fmt"
	"math"
)

// func ReadFile(filename string) (string, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()
// 	rd := bufio.NewReader(file)
// 	line, _ := rd.ReadString('\n')
// 	return line, nil
// }

// func WirteFile(filename string, line string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	_, err = fmt.Fprintln(file, line)
// 	return err
// }

// const Filename string = "data.txt"

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func SqrtExe() {
	sqrt, err := sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func registerAccount(name, passward string) error {
	if len(passward) < 8 {
		return PasswordError{len(passward), 0}
	}
	return nil
}

func RegisterAccountPrint() {
	err := registerAccount("my10", "myPw")
	if err == nil {
		fmt.Println("회원 가입됐습니다.")
	}
	if errInfo, ok := err.(PasswordError); ok {
		fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
	}
}

// func MultipleFromString(str string) (int, error) {
// 	scanner := bufio.NewScanner(strings.NewReader(str))
// 	scanner.Split(bufio.ScanWords)

// 	pos := 0
// 	a, n, err := readNextInt(scanner)
// 	if err != nil {
// 		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
// 	}
// 	pos += n + 1
// 	b, n, err := readNextInt(scanner)
// 	if err != nil {
// 		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
// 	}
// 	return a * b, nil
// }

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()
	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func PanicFunc() {
	f()
	fmt.Println("프로그램이 계속 실행됨")
}
