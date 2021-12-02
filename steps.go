package main

import (
	"fmt"
	"log"
	"math"
)

func PlaceFood() {

	for i, f := range Food {
		if f.X == -1 {
			for {
				j := randInt(Program.WorldSize)
				if World[j] == 0 {
					World[j] = -2
					p := IndexToXY(j)
					Food[i] = p
				}
			}
		}
	}
}

// Step1 : Fill in all the inputs from the environment.
func Step0(i int) {

	fWorldY := float64(Program.WorldY)
	fWorldX := float64(Program.WorldX)
	nLocX := float64(Neos[i].LocationX)
	nLocY := float64(Neos[i].LocationY)
	fCurrentStep := float64(CurrentStep)
	fNumberOfSteps := float64(Program.NumberOfSteps)

	//distanceFromNorth := fWorldY - nLocY/fWorldY*2.0 - 1.0
	//distanceFromWest := fWorldX - nLocX/fWorldX*2.0 - 1.0

	distanceFromNorth := nLocY/fWorldY*2.0 - 1.0 // distanceFromNorth : Distance from the North wall. -1.0 to 1.0. -1 is full north, 1 is full south
	distanceFromEast := nLocX/fWorldX*2.0 - 1.0  // distanceFromWest : Distace from the East Wall. -1.0 to 1.0. -1 is full east, 1 is full west

	fMaxDist := float64(Program.MaxDistanceLook)

	Neos[i].Inputs[0] = fCurrentStep / fNumberOfSteps
	Neos[i].Inputs[3] = distanceFromNorth
	Neos[i].Inputs[4] = distanceFromEast
	Neos[i].Inputs[7] = float64(Neos[i].Hunger) / float64(Program.MaxHunger)

	fmt.Println(i, fWorldX, nLocX, distanceFromEast)

	// Blockage forward
	for indexS := 1; indexS < Program.MaxDistanceLook; indexS++ { // going to look to far at anlges but will fix at distance check

		p := DirectionToStep(Neos[i].Direction)

		var td float64 // td : Temporary distance to a target.
		tx := Neos[i].LocationX + (indexS * p.X)
		ty := Neos[i].LocationY + (indexS * p.Y)

		ftx := float64(tx)
		fty := float64(ty)

		// Check for the end of world
		if tx >= Program.WorldX || ty >= Program.WorldY || tx < 0 || ty < 0 {
			if tx > Program.WorldX {
				tx = Program.WorldX
			}
			if ty > Program.WorldY {
				ty = Program.WorldY
			}
			if tx < 0 {
				tx = 0
			}
			if ty < 0 {
				ty = 0
			}
			td = GetDistance(nLocX, nLocY, ftx, fty)
			if td > fMaxDist {
				Neos[i].Inputs[8] = -1.0
				break
			}
			Neos[i].Inputs[8] = td
			break
		}

		if tx >= Program.WorldX || ty >= Program.WorldY {
			log.Fatalln(Neos[i].LocationX, Neos[i].LocationY, "tx", tx, "ty", ty)
		}

		if World[XYtoIndex(tx, ty)] != 0 {
			td = GetDistance(nLocX, nLocY, ftx, fty)
			if td > fMaxDist {
				Neos[i].Inputs[8] = -1.0
				break
			} else {
				Neos[i].Inputs[8] = td / fMaxDist
				break
			}
		}
		if indexS == Program.MaxDistanceLook-1 {
			Neos[i].Inputs[8] = -1.0
		}

	}

	// Blockage backwards
	for indexS := 1; indexS < Program.MaxDistanceLook; indexS++ { // going to look to far at anlges but will fix at distance check

		db := Neos[i].Direction
		if db < 180 {
			db += 180
		} else {
			db -= 180
		}

		p := DirectionToStep(db)

		var td float64 // td : Temporary distance to a target.
		tx := Neos[i].LocationX + (indexS * p.X)
		ty := Neos[i].LocationY + (indexS * p.Y)

		ftx := float64(tx)
		fty := float64(ty)

		// Check end of world
		if tx >= Program.WorldX || ty >= Program.WorldY || tx < 0 || ty < 0 {
			if tx > Program.WorldX {
				tx = Program.WorldX
			}
			if ty > Program.WorldY {
				ty = Program.WorldY
			}
			if tx < 0 {
				tx = 0
			}
			if ty < 0 {
				ty = 0
			}
			td = GetDistance(nLocX, nLocY, ftx, fty)
			if td > fMaxDist {
				Neos[i].Inputs[9] = -1.0
				break
			}
			Neos[i].Inputs[9] = td
			break
		}

		if World[XYtoIndex(tx, ty)] != 0 {
			td = GetDistance(nLocX, nLocY, ftx, fty)
			if td > fMaxDist {
				Neos[i].Inputs[9] = -1.0
				break
			} else {
				Neos[i].Inputs[9] = td / fMaxDist
				break
			}
		}

		if indexS == Program.MaxDistanceLook-1 {
			Neos[i].Inputs[9] = -1.0
		}

	}

	// -----------------------------------

	d := 100000.0
	a := 0.0
	for _, k := range Food {
		fkx := float64(k.X)
		fky := float64(k.Y)
		d1 := GetDistance(nLocX, nLocY, fkx, fky)
		a1 := GetAngle(nLocX, nLocY, fkx, fky)
		if d1 < d {
			d = d1
			a = a1
		}
	}
	if d <= fMaxDist {
		Neos[i].Inputs[5] = d / fMaxDist
		Neos[i].Inputs[1] = a / 360.0
	} else {
		Neos[i].Inputs[5] = -1.0
		Neos[i].Inputs[1] = -1.0
	}

	d = 100000.0
	a = 0.0
	for j, k := range Neos {
		if j == 0 { // Skip 0
			continue
		}
		if j == i {
			continue
		}
		fkLocX := float64(k.LocationX)
		fkLocY := float64(k.LocationY)
		d1 := GetDistance(nLocX, nLocY, fkLocX, fkLocY)
		a1 := GetAngle(nLocX, nLocY, fkLocX, fkLocY)
		if d1 < d {
			d = d1
			a = a1
		}
	}
	if d <= fMaxDist {
		Neos[i].Inputs[6] = d / fMaxDist
		Neos[i].Inputs[2] = a / 360.0
	} else {
		Neos[i].Inputs[6] = -1.0
		Neos[i].Inputs[2] = -1.0
	}

}

// Step1 : Clear all neuron's inputs and move the Neo's env inputs to the neurons.
func Step1(i int) {

	for j, n := range Neos[i].Neurons {
		Neos[i].Neurons[j].InValue = 0.0
		if n.SourceLayer == 0 {
			Neos[i].Neurons[j].InValue = Neos[i].Inputs[n.Source]
		}
	}

}

// Step2 : Propagate out to in and sum and pass through Tanh.
func Step2(i int) {

	if i > Program.NumberOfNeos || i < 1 {
		log.Fatalf("Step2 id '%d' is out of bounds", i)
	}

	for j := 0; j < Program.NumberOfLayers-1; j++ { // Layer by Layer - why -1 ?

		for k, m := range Neos[i].Neurons { // Loop through the Neurons

			if m.SourceLayer == j { // Does SourceLayer match loop j

				Neos[i].Neurons[k].OutValue = m.InValue * m.Weight

				for n, o := range Neos[i].Neurons { // Loop through the neurons again looking for match

					if n == k { // Skip if same neuron
						continue
					}

					if o.SourceLayer == m.OutLayer && o.Out == m.Out {
						Neos[i].Neurons[n].InValue += Neos[i].Neurons[k].OutValue
					}

				}

			}

		}

		if j != 0 {
			for k, m := range Neos[i].Neurons {
				if m.SourceLayer == j {
					// shouldnt this be outvalue?
					//Neos[i].Neurons[k].InValue = math.Tanh(m.InValue)
					Neos[i].Neurons[k].OutValue = math.Tanh(m.InValue)
				}
				if m.OutLayer == Program.NumberOfLayers-1 { // put the output layer into the neo's output slilce.
					Neos[i].Outputs[m.Out] = m.OutValue
				}
			}
		}
	}

}

// Step3 : Check if Neos died from hunger.
func Step3(i int) {

	if Neos[i].Dead {
		return
	}

	// Check if Neo died from hunger
	if Neos[i].Hunger == 0 {
		Neos[i].Dead = true
		return
	}
}

// Step4 : Do movement
func Step4(i int) {

	var sX float64
	var sY float64

	sX += Neos[i].Outputs[7]
	sY += Neos[i].Outputs[6]

}

// probability : return true if random is less than p.
func probability(p float64) bool {
	if randFloat() <= math.Abs(p) {
		return true
	}
	return false
}
