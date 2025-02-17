package oops

import (
	"fmt"
)

// --- Encapsulation ---
// Define a struct to represent a Person with encapsulated fields.
// Unexported fields (private) start with a lowercase letter.
type Person struct {
	name string // Unexported field (private)
	age  int    // Unexported field (private)
}

// Exported method (public) to access the private name field.
func (p Person) GetName() string {
	return p.name
}

// Exported method (public) to access the private age field.
func (p Person) GetAge() int {
	return p.age
}

// Constructor-like function to create a new Person.
func NewPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

// --- Composition ---
// Define a struct to represent an Employee, which embeds the Person struct.
// This is an example of composition (reusing fields and methods from Person).
type Employee struct {
	Person  // Embedding Person struct
	company string
}

// Method specific to Employee.
func (e Employee) Work() string {
	return fmt.Sprintf("%s is working at %s", e.GetName(), e.company)
}

// Constructor-like function to create a new Employee.
func NewEmployee(name string, age int, company string) Employee {
	return Employee{
		Person:  NewPerson(name, age), // Reuse Person constructor
		company: company,
	}
}

// --- Polymorphism ---
// Define an interface called Speaker.
// Any type that implements the Speak() method satisfies this interface.
type Speaker interface {
	Speak() string
}

// Define a struct to represent a Dog.
type Dog struct {
	name string
}

// Implement the Speak() method for Dog.
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.name)
}

// Define a struct to represent a Cat.
type Cat struct {
	name string
}

// Implement the Speak() method for Cat.
func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.name)
}

// Function to demonstrate polymorphism.
// It accepts any type that satisfies the Speaker interface.
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

// --- Abstraction ---
// Define an interface to represent a Vehicle.
// This abstracts the behavior of starting a vehicle.
type Vehicle interface {
	Start() string
}

// Define a struct to represent a Car.
type Car struct {
	model string
}

// Implement the Start() method for Car.
func (c Car) Start() string {
	return fmt.Sprintf("%s is starting... Vroom!", c.model)
}

// Function to demonstrate abstraction.
// It accepts any type that satisfies the Vehicle interface.
func Drive(v Vehicle) {
	fmt.Println(v.Start())
}

func main() {
	// --- Encapsulation Example ---
	person := NewPerson("Alice", 30)
	fmt.Println("Person Name:", person.GetName()) // Access private field via method
	fmt.Println("Person Age:", person.GetAge())   // Access private field via method

	// --- Composition Example ---
	employee := NewEmployee("Bob", 25, "TechCorp")
	fmt.Println(employee.Work())                      // Access method from Employee
	fmt.Println("Employee Name:", employee.GetName()) // Access method from embedded Person

	// --- Polymorphism Example ---
	dog := Dog{name: "Buddy"}
	cat := Cat{name: "Whiskers"}
	MakeSound(dog) // Buddy says: Woof!
	MakeSound(cat) // Whiskers says: Meow!

	// --- Abstraction Example ---
	car := Car{model: "Tesla Model S"}
	Drive(car) // Tesla Model S is starting... Vroom!
}
