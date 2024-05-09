package main
import "fmt"
func main(){
	num :=[]int{2,3,4}
	sum:=0
	for _, num:= range num{ //index is not needed, so _ --> blank identifier
		sum+=num
	}
	fmt.Println("sum: ",sum)

	for i, num:= range num { // i stores the index of elements of slice
		if num ==3{
			fmt.Println("index: ",i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k, v:= range kvs{ //k gives key and v gives value
		fmt.Printf("%s->%s\n",k,v)
	}

	for k := range kvs{
		fmt.Println("key: ",k)
	}
	for i, c:=range "go"{ //range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
		fmt.Println(i,c)
	}
}