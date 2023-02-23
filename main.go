package main

import "fmt"

type Profile struct {
	Name string
	ImageURL string
	Age int
	Gender string
	Description string
}

type Evidence struct {
	Name string	
	ImageURL string
	Type string
	Description string
	SubmitPerson string
}

func (p Profile) Check() {
	fmt.Printf("--Profile--\nName: %v\nAge: %v\nGender: %v\nDescription: %v\n", p.Name, p.Age, p.Gender, p.Description)
}

func (e Evidence) Check() {

}

func (e Evidence) Present() {

}

func main() {
	gumshoe := Profile{
		Name: "Dick Gumshoe",
		ImageURL: "",
		Age: 30,
		Gender: "Male",
		Description: "Detective at the local precinct, In charge of the initial investigation.",
	}

	gumshoe.Check()
}