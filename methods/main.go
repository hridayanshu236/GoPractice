package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	hridayanshu := User{"Hridayanshu","hridayanshu23@gmail.com",true, 20}
	fmt.Println(hridayanshu)
	fmt.Printf("Hridayanshu's details are: %+v\n", hridayanshu)
	fmt.Printf("Name is: %v\n", hridayanshu.Name)
	hridayanshu.GetStatus()
	hridayanshu.newMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ",u.Status)
}

func (u User) newMail(){
	u.Email = "test@go.dev" //doesnt changes the original email, it just manipulates and prints
	fmt.Println("Email of this user is: ",u.Email)
}
