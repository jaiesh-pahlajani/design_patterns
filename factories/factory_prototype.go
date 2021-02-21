package main

import "fmt"

type Employee2 struct {
	Name, Position string
	AnnualIncome int
}

const (
	Developer = iota
	Manager
)

// functional
func NewEmployee(role int) *Employee2 {
	switch role {
	case Developer:
		return &Employee2{"", "Developer", 60000}
	case Manager:
		return &Employee2{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}