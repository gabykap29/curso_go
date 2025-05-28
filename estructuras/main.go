package main

import "fmt"

type User struct {
	name  string
	email string
}

func (self *User) setName(name string) {
	self.name = name
}
func (self *User) setEmail(email string) {
	self.email = email
}
func main() {
	var gab User

	gab.name = "Gabriel"
	gab.email = "example@gmail.com"

	mar := User{
		name:  "Mariela",
		email: "exampleQgmail.com",
	}
	name := "Ethan"
	email := "example@gmail.com"

	eth := User{name, email}

	emm := User{
		name:  "",
		email: "",
	}

	emm.setName(("Emma"))
	emm.setEmail("example@gmail.com")
	fmt.Println(gab)
	fmt.Println(mar)
	fmt.Println(eth)
	fmt.Println(emm)
}
