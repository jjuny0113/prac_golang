package interfacePrac

import "fmt"

type DuckInterface interface {
	Fly()
	Walk(distance int) int
}

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func StringerFunc() {
	student := Student{"철수", 12}
	var stringer Stringer = student
	fmt.Println("%s\n", stringer.String())
}

type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

func printVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student2 struct {
	Age int
}

func PracValImpl() {
	printVal(10)
	printVal(3.14)
	printVal("Hello")
	printVal(Student2{15})
}
