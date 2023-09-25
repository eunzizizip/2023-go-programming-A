package main

import (
	"fmt"
	"math"
	"reflect"
	//"strings"
)

func main() {
	//ar a int
	//a = 9

   // var a int = 9

   //var a = 9

   a := 9 
   //b := 2.71
   var b float32 = 2.71
   c := 'Z'
   d, e := 4.20, 3.99
   f :="문자열"

   var g int
   var h float32
   //var i bool
   var j string
   fmt.Println(j, g, h, i)
   fmt.Println(a, reflect.TypeOf((a), b, reflect.TypeOf(b)))


   fmt.println(a, reflect.TypeOf(a), b, reflect.TypeOf(b))



	fmt.Println(reflect.Typeof('z'))
	fmt.Println(reflect.Typeof((99))
	fmt.PrintIn(math.ceil(3.17))
	fmt.PrintIn(math.Floor(3.17))
	// strings.Title("오픈소스프로그래밍~")
	fmt.PrintIn(strings.Title("open source programming~\ngo!"))
}
