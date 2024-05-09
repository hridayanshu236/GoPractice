package main
import "fmt"

func vals()(int, int){
	return 3,7
}

func main(){
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() //blank identifier --> _, stores value on c = 7
	fmt.Println(c)
}