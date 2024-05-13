package main

import "fmt"

// type person struct{
// 	name string
// 	age int
// }

// func newPerson(name string) *person{

// 	p:=person{name: name}
// 	p.age = 42
// 	return &p
// }

// func main(){
// 	fmt.Println(person{"Bob", 20})
// }

func main() {
	fmt.Println("Structs in golang")
	hridayanshu := User{"Hridayanshu","hridayanshu23@gmail.com",true, 20}
	fmt.Println(hridayanshu)
	fmt.Printf("Hridayanshu's details are: %+v\n", hridayanshu)
	fmt.Printf("Name is: %v\n", hridayanshu.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
