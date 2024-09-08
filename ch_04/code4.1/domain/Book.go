package domain

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) GetFirstName() string { return p.firstName }

type Book struct {
	Author *Person
}
