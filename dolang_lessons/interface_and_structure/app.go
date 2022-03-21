package main

import (
	"fmt"
)

func main() {

	var fork TwoSidesPrimitive
	fork = NewSampleStruct(2, 3)
	fmt.Println(fork.SumRoute())
	fmt.Printf("%v %T", fork, fork)
}

type SampleStruct struct {
	left, right int
}

func NewSampleStruct(left, right int) TwoSidesPrimitive {
	return &SampleStruct{left: left, right: right}
}

type TwoSidesPrimitive interface {
	MaxDirVal() int
	MaxDirRoute() string
	SumRoute() int
}

func (receiver SampleStruct) MaxDirVal() int {
	return 0
}
func (receiver SampleStruct) SumRoute() int {
	return 0
}

func (receiver SampleStruct) MaxDirRoute() string {
	return "left"
}
