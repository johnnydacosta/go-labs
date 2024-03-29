package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("%s %d", u.Name, u.Age)
}

func (u User) ItMyBirthday() {
	u.Age++
	fmt.Printf("Hey %s, you just turned %d\n", u.Name, u.Age)
}

func (u *User) ItMyBirthdayP() {
	u.Age++
	fmt.Printf("Hey %s, you just turned %d\n", u.Name, u.Age)
}

func main() {
	user := User{Name: "Johnny", Age: 31}
	fmt.Printf("Hello %s\n", user)
	user.ItMyBirthday()
	fmt.Println(user)
	user.ItMyBirthdayP()
	fmt.Println(user)
}
