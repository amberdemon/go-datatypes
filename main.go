package main

import (
	"fmt"

	"github.com/amberdemon/go-datatypes/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropenUnionIdentifier("123-45-678", "Germany"))
	err := p.SetTwitterHandler("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.Country())
	println(p.FullName())
}
