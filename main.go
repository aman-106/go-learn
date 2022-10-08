package main

import "fmt"
import "errors"

const (
	first = iota + 8
	second 
	third 
)

const (
	fourth = iota + 4  // complex const expressions , dynamic
)

func main(){  
	fmt.Println("hello woels")

	var i int = 2
	var f float32 = 3.34 // value types  - poin to value 
	b := true
	fmt.Println(i,f,b)
	// pointer - point to address of value

	var str *string = new(string)  // memory alloc for pointer
	*str = "Author"
	fmt.Println(*str) // no pointer aithemetic ( str + 1 )
	fmt.Println(str)

	 str2 := "new Author"
	 ptr2 := &str2;

	 fmt.Println(str2,ptr2)

	 // iota 

	fmt.Println(first,second,fourth)

	// array 
	var arr [3]int 
	arr[0] = 1
	arr[1] = 2
	arr[2] = 10
	// arr[3] = 10 bounded 
	 arr2 := [3]int{1,2,3} 
	 slice := arr2[:]
	fmt.Println(arr,slice);

	// square brackets
	slice2 := []int{23,383,3838} //resizing array , alloc memory dynamic 
	fmt.Println(slice2); 

	m := map[string]string{};
	m["sec"] = "9399"
	m["se2c"] = "9399"

	fmt.Println(m);

	 type user struct {
		 id int
		 name string
	 }

	 p := user{id:10,name:"aman"}
	 fmt.Println(p)

	//  func newPerson() *user{
	// 	 return &p;
	//  
	po , err := createNewServer(10)
	fmt.Println(po,err)

	map2 := map[string]int {"ssksk":1929,"djd":1939}

	for _, v := range map2 {
		println(v)
	}

	// cliApp();
	argsReader();
	getHelp()
}


func createNewServer(port int) (int,error)  {

	return 100,errors.New("type")
} 
  