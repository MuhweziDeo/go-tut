package main

import ("fmt")

type Student struct {
	name string
	age int
	grades []int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() int {
	sum := 0

	for _, v := range s.grades {
		sum += v
	}
	return sum /len(s.grades)
}


func main() {
	s := Student{"deo", 27, []int{90, 90, 40}}

	s.setAge(29)
	fmt.Println(s.getAge())

	fmt.Println(s.getAverageGrade())

}