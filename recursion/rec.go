package main 
import "fmt"

func fact(n int) int{
	if n==0{
		return 1
	}
	return n*fact(n-1)
}

func main(){
	fmt.Println(fact(7))
	var fib func(n int) int //closures can also be recursive, but this requires the closure to be declared with a type var explicitly before its defined.
	 fib = func(n int) int{
		if n<2{
			return n
		}
		return fib(n-1)+fib(n-2)
	 }

	 fmt.Println(fib(7))
}