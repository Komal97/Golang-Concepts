package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}
func implementStruct(){

	fmt.Println("\n--------------- STRUCT -----------------")

	// method - 1
	// alex1 := person{"Alex", "Anderson"}

	// method - 2
	alex2 := person{firstName: "Alex",
					lastName: "Anderson",
					contact: contactInfo{
						email: "alex123@gmail.com",
						zipCode: 94000,
					},
			}
	fmt.Println(alex2)

	// method - 3
	var alex3 person
	alex3.firstName = "Alex"
	alex3.lastName = "Anderson"
	alex3.contact.email = "alex123@gmail.com"
	alex3.contact.zipCode = 94000

	alex3.printObj()

	//ptrAlex := &alex3
	//ptrAlex.updateName("Alexa")

	// shortcut
	alex3.updateName("Alexa")
	alex3.printObj()

}

//func (p person) updateName(newFirstName string){
//	p.firstName = newFirstName
//}

func (ptrPerson *person) updateName(newFirstName string){
	(*ptrPerson).firstName = newFirstName
}

func (p person) printObj(){
	fmt.Printf("%+v", p)
	fmt.Println()
}