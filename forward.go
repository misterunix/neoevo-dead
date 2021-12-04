package main

import "fmt"

// step0 : Clear inputs.
func step0(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

	for i := range Neos[id].Neurons {
		Neos[id].Neurons[i].InValue = 0.0
		Neos[id].Neurons[i].OutLayer = 0.0
	}

	return nil

}

// step1 : Move env to inputs.
func step1(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

}
