package main
import "fmt"

var a int = 2

func main(){
	
//	fmt.Println(a)
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "Bu Tsikl")
		if i == a{
			
			fmt.Println("->", a, "Bu if")
			a += 2

		}
	}
}