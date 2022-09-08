package user
// package main
// import "fmt"
// var l = fmt.Println

var users =  make(map[int]string)
var id int = 0

func Create(name string) int {
	id++
	users[id] = name
	return id
}

func Del(id int) bool{
	_,ok := users[id]
	if ok {
		delete(users,id)
		return true
	}
	return false
}

func Update(id int,name string) bool{
	_,ok := users[id]
	if ok {
		users[id] = name
		return true
	}
	return false
}

func Get(id int) string{
	data,ok := users[id]
	if ok {
		return data
	}
	return ""
}

// func main()  {
// 	Create("zs")
// 	l(Get(1))
// }