package main

import (
	"fmt"
//	"math"
//	"time"
)

func main(){
	testif()
	testif2()
}

func testif(){
	i := 8
//	i = abs(-8)
	if i % 2== 0 {
		fmt.Println("i Juft son")
	} else if i % 2 != 0 {
		fmt.Println("i Juft son emas")
	} 
}

func testif2(){
	k := 4.0
	var t float64 = k*4
	if t == 16.0{
		fmt.Println("K soni 4ga teng")
	} else {fmt.Println("K soni 4dan boshqa son")}

}