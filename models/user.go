package models

import "fmt"

type User struct {
	Id int
	Name string
}

var (
	users []*User
	currentUserIndex int = 1
)


func appendUser(name string){
	u := User{Id:currentUserIndex,Name:name}
	users = append(users,&u);
	// append()
	currentUserIndex++;
	fmt.Println(users);
}


func add(p int) int{
	return p+1;
}

