package main

import (
	"fmt"

	"github.com/YarramallaAjay/GO/designpatterns/builderpattern"
)

func main() {
	builder := &builderpattern.ConcretePaperBoat{}
	maker := &builderpattern.BoatMaker{}

	boat := maker.BoatMakerBuilder(
		80,    // paper size
		true,  // cryons
		false, // ruler
		true,  // water tub
		builder,
	)

	fmt.Println("Paper Boat Built:")
	fmt.Printf("%+v\n", boat)
}
