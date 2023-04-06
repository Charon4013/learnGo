package main

import "fmt"

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	level int
	user
}

func (a admin) notify() {
	fmt.Printf("admin %s notified\n", a.name)
}

// 值接收
func (u user) notify() {
	fmt.Printf("user %s notified\n", u.name)
}

// 值接收
func (u user) printName() {

}

// 指针接收
func (u *user) changeName(newName string) {
	u.name = newName
}

func main() {
	u := user{
		name:  "Pan",
		email: "pan@qq.com",
	}

	a := admin{
		level: 0,
		user: user{
			name:  "Chen",
			email: "chen@qq.com",
		},
	}

	sendNotification(a)
	sendNotification(u)
}
