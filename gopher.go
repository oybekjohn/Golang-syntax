package main
import "fmt"
var s int = 10
var m int = 15

var n = 2
var k = n

func main(){
	fmt.Println("gopher")
	fmt.Print(s+m)
	fmt.Println(k)
	for i := 1; i <= 3; i++ {
		fmt.Print(i)
	}
}