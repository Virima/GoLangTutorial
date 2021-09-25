package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// function named getAge() that return int
func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	s1 := Student{"Kelvin", []int{70, 90, 80, 85}, 17}

	fmt.Println(s1.getAge())
	s1.setAge(28)
	fmt.Println(s1.getAge())

	average := s1.getAverageGrade()
	fmt.Println(average)

	maxVal := s1.getMaxGrade()
	fmt.Println(maxVal)
}
