package models

type Person struct {
	ID       int
	Name     string
	LastName string
}

var Persons = []Person{
	{ID: 1, Name: "Ivan", LastName: "Ivanov"},
	{ID: 2, Name: "Maria", LastName: "Petrova"},
}
