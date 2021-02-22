package main

import "fmt"

type Address2 struct {
	StreetAddress, City, Country string
}

func (a *Address2) DeepCopy() *Address2 {
	return &Address2{
		a.StreetAddress,
		a.City,
		a.Country }
}

type Person2 struct {
	Name string
	Address *Address2
	Friends []string
}

func (p *Person2) DeepCopy() *Person2 {
	q := *p // copies Name
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	john := Person2{"John",
		&Address2{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}


