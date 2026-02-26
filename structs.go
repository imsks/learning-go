package main

func learnStructs() {
	// 1. Defining and creating structs

	// type Person struct {
	// 	Name string
	// 	Age  int
	// }

	// p1 := Person{
	// 	Name: "Sachin",
	// 	Age:  27,
	// }

	// p2 := Person{"Sachin1", 27}

	// var p3 Person
	// p3.Name = "Sachin2"
	// p3.Age = 27

	// p4 := new(Person)
	// p4.Name = "Sachin3"

	// fmt.Println(p1, p2, p3, p4)

	// // 2. Struct fields and access
	// type Employee struct {
	// 	ID         int
	// 	Name       string
	// 	Department string
	// 	Salary     float64
	// }

	// emp := Employee{
	// 	ID:         1,
	// 	Name:       "Sachin",
	// 	Department: "Engineering",
	// 	Salary:     5000,
	// }

	// fmt.Println(emp.Salary)
	// emp.Salary = 6000
	// fmt.Println(emp.Salary)

	// // 3. Methods – value receiver
	// type Person struct {
	// 	Name string
	// 	Age  int
	// }

	// func (p Person) Greet() string {
	// 	return "Hello, I'm " + p.Name
	// }

	// // 6. Struct embedding (composition)
	type Address struct {
		City  string
		State string
	}

	type Person struct {
		Name string
		Age  int
		Address
	}

	// p := Person{
	// 	Name: "Sachin",
	// 	Age:  25,
	// 	Address: Address{
	// 		City:  "Lakhimpur Kheri",
	// 		State: "Uttar Pradesh",
	// 	},
	// }

	// fmt.Println(p.Address.City)

	// // 7. Pointers to structs
	// p := &Person{Name: "Sachin", Age: 27}
	// fmt.Println(p.Name) // (*p).Name = p.Name as Go automatically dereferences for fields

	// type Book struct {
	// 	Title string
	// }

	// func Summary(name string) *Book {
	// 	return &Book{
	// 		Title: "Book Title",
	// 	}
	// }
}
