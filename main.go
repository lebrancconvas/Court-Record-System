package main

import "fmt"

type Data interface {
	Check()
}

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
	Source string
}

type Testimony struct {
	Declarator string
	Statement []string
}

func (p Profile) Check() {
	fmt.Printf("--Profile--\nName: %v\nAge: %v\nGender: %v\nDescription: %v\n", p.Name, p.Age, p.Gender, p.Description)
}

func (e Evidence) Check() {
	fmt.Printf("--Evidence--\nName: %v\nType: %v\nDescription: %v\nSource: %v\n", e.Name, e.Type, e.Description, e.Source)
}

func (e Evidence) Present() {
	fmt.Println("Take That!!")
	fmt.Printf("Displaying Evidence: %v\n", e.Name)
}

func main() {
	gumshoe := Profile{
		Name: "Dick Gumshoe",
		ImageURL: "https://static.wikia.nocookie.net/aceattorney/images/6/65/Sprite-gumshoe.gif/revision/latest/scale-to-width-down/644?cb=20120226001744",
		Age: 30,
		Gender: "Male",
		Description: "Detective at the local precinct, In charge of the initial investigation.",
	}

	gumshoe.Check()
}