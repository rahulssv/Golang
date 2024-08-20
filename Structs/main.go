package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var rahul person
	// rahul.firstName="Rahul"
	// rahul.lastName="Vishwakarma"
	// rahul.contact= contactInfo{email: "rahul@gmail.com", zipcode: 403401,}
	rahul = person{
		firstName: "Rahul",
		lastName:  "Vishwakarma",
		contact: contactInfo{
			email:   "rahul@gmail.com",
			zipcode: 403401,
		},
	}
	fmt.Println(*&rahul.firstName)
	rahul.updateName("Rohit")
	rahul.print()

}
func (p *person) updateName(newFirstName string){
	p.firstName = newFirstName
}

func (p person) print(){
	fmt.Println(p)
}