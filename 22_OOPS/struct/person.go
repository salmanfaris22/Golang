// person.go
package person

// Person struct with unexported fields
type Person struct {
	name string
	age  int
}

// NewPerson is an exported function to create a new Person
func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

// GetName is an exported method to get the person's name
func (p *Person) GetName() string {
	return p.name
}

// SetName is an exported method to set the person's name
func (p *Person) SetName(name string) {
	p.name = name
}

// GetAge is an exported method to get the person's age
func (p *Person) GetAge() int {
	return p.age
}

// SetAge is an exported method to set the person's age
func (p *Person) SetAge(age int) {
	p.age = age
}
