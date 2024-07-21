package slice

import (
	"fmt"
	"sort"
)

func Slice() {
	var slice = []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		slice[i] += 10
	}
	for i, v := range slice {
		slice[i] = v * 2
	}
	slice = append(slice, 77)
	fmt.Println(slice)
	// sliceing := slice[1:3]
	// sliceing2 := slice[1:]
	// sliceing3 := slice[:]
	// fmt.Println(sliceing)
	// fmt.Println(sliceing2)
	// fmt.Println(sliceing3)
	copyExam1 := []int{1, 2, 3, 4, 5}
	copyExam2 := make([]int, 3, 10)
	copyExam3 := make([]int, 10)
	cnt := copy(copyExam2, copyExam1)
	fmt.Println(copyExam1)
	fmt.Println(copyExam2)
	fmt.Println(copyExam3)
	fmt.Println(cnt)
}

func Pop() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2
	// for i := idx + 1; i < len(slice); i++ {
	// 	slice[i-1] = slice[i]
	// }
	// slice = slice[:len(slice)-1]
	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Println(slice)
}

func Add() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2
	// slice = append(slice, 0)
	// for i := len(slice) - 2; i >= idx; i-- {
	// 	slice[i+1] = slice[i]
	// }
	// slice[idx] = 100
	// slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100
	fmt.Println(slice)
}

func Sort() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}

type Student struct {
	Name string
	Age  int
}

type Students []Student

func SortStruct() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42}, {"켄", 38}, {"송하나", 18},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
}

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
