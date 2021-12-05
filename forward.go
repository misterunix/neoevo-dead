package main

import (
	"fmt"
	"math"
)

// forward0 : Clear inputs.
func forward0(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

	for i := range Neos[id].Neurons {
		Neos[id].Neurons[i].InValue = 0.0
		Neos[id].Neurons[i].OutValue = 0.0
	}

	return nil

}

// forward1 : Move env to inputs
func forward1(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

	// Age
	Neos[id].Inputs[AGE] = float64(step) / float64(Program.NumberOfSteps)

	Neos[id].Inputs[POSITIONNS] = ((Neos[id].Location.FY / Program.FworldY) * 2.0) - 1.0
	Neos[id].Inputs[POSITIONWE] = ((Neos[id].Location.FX / Program.FworldX) * 2.0) - 1.0

	return nil
}

// forward2 : Move inputs into layer 0 neurons
func forward2(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

	for i, n := range Neos[id].Neurons {
		if n.SourceLayer == 0 {
			Neos[id].Neurons[i].InValue = Neos[id].Inputs[n.SourceID] * n.Weight
			Neos[id].Neurons[i].OutValue = Neos[id].Inputs[n.SourceID]
		}
	}

	return nil

}

// forward3 : propaget forward
func forward3(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

	for l := 0; l < Program.NumberOfLayers; l++ {

		for _, n := range Neos[id].Neurons {

			if n.SourceLayer == l {

				if n.LinkForward >= 0 {
					Neos[id].Neurons[n.LinkForward].InValue += (n.OutValue * n.Weight)
					fmt.Println(Neos[id].Neurons[n.LinkForward].InValue)
					fmt.Println(Neos[id].Neurons[n.LinkForward].InValue, n.OutValue, n.Weight)
				}
				if n.LinkForward == -99 {
					Neos[id].Outputs[n.OutID] += n.OutValue
				}

			}

		}

		for i, n := range Neos[id].Neurons {

			if n.SourceLayer == 0 {
				Neos[id].Neurons[i].OutValue = n.InValue
			} else if n.SourceLayer == Program.NumberOfLayers-1 {
				Neos[id].Outputs[n.OutID] = math.Tanh(Neos[id].Outputs[n.OutID])
			} else {
				Neos[id].Neurons[i].OutValue = math.Tanh(n.InValue)
			}

		}

	}

	return nil

}

// forward4 : move out neurons to neo output
func forward4(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

	for _, n := range Neos[id].Neurons {
		if n.OutLayer == Program.NumberOfLayers-1 {
			Neos[id].Outputs[n.OutID] = n.InValue // math.Tanh(n.InValue * n.Weight)
		}
	}

	return nil

}
