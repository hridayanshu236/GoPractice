package main
import(
	"fmt"
	"time"
)

func main(){
	i:=2
	fmt.Println("Write", i, "as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){ //weekday is used to check the name of the day
		case time.Saturday, time.Sunday:
			fmt.Println("It is weekend")
		default:
			fmt.Println("Its a weekday")
	}

	t:= time.Now() //gives the current time
	switch{
	case t.Hour() <12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	whatAmI := func(i interface{}){ //empty interface holds values of any type
		switch t:=i.(type){
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Dont know type %T\n",t) //println adds new line at the end
												//while printf doesnt
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")


}