package main

import(
	"fmt"
)

func sum(x, y int) (int, string){
	return x + y, "string as well"

}

type rectangle struct{
	length int
	breadth int
}
func (r rectangle) dimension() (int, int){
	return r.length, r.breadth;
}
func (r rectangle) area() int{
	return r.length * r.breadth
}
func main() {
	fmt.Println("heyoo")

	var name string = "JIIT"

	fmt.Printf("welcome to %v\n", name)

	var(
		i int  = 2
		//f float64 = 3.14
	)
	fmt.Printf("i: %v with type %T\n", i, i)

	flag := true
	fmt.Printf("flag is :%v of type %T\n", flag, flag)

	fmt.Println(sum(2,3))
	r := rectangle{10,20}
	fmt.Println(r.dimension())
	fmt.Println(r.area())

}
