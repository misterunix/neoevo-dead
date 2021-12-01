package main

import (
	"log"
	"math"
	"sync"
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

// Step1 : cycle through all Neos and fill the inputs
func Step0() {

	for i := range Neos {
		if i == 0 { // start with 1
			continue
		}
		distanceFromNorth := float64(Program.WorldY-Neos[i].LocationY)/float64(Program.WorldY)*2.0 - 1.0
		distanceFromWest := float64(Program.WorldX-Neos[i].LocationX)/float64(Program.WorldX)*2.0 - 1.0
		Neos[i].Inputs[0] = float64(CurrentStep) / float64(Program.NumberOfSteps)
		Neos[i].Inputs[3] = distanceFromNorth
		Neos[i].Inputs[4] = distanceFromWest
		Neos[i].Inputs[7] = float64(Neos[i].Hunger) / float64(Program.MaxHunger)

		// Blockage forward
		// akin to raytracing, which I hate
		// simple populate world array with items them step the direction
		for indexS := 1; indexS < Program.MaxDistanceLook; indexS++ { // going to look to far at anlges but will fix at distance check

			p := DirectionToStep(Neos[i].Direction)

			var td float64 // td : temporary distance to a target
			tx := Neos[i].LocationX + (indexS * p.X)
			ty := Neos[i].LocationY + (indexS * p.Y)

			// Check end of world
			if tx >= Program.WorldX || ty >= Program.WorldY || tx < 0 || ty < 0 {
				td = GetDistance(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(tx), float64(ty))
				if td > float64(Program.MaxDistanceLook) {
					Neos[i].Inputs[8] = -1.0
					//fmt.Println("Should break", tx, ty)
					break
				}
				Neos[i].Inputs[8] = td
				break
			}

			if tx >= Program.WorldX || ty >= Program.WorldY {
				log.Fatalln(Neos[i].LocationX, Neos[i].LocationY, "tx", tx, "ty", ty)
			}

			if World[XYtoIndex(tx, ty)] != 0 {
				td = GetDistance(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(tx), float64(ty))
				if td > float64(Program.MaxDistanceLook) {
					Neos[i].Inputs[8] = -1.0
					break
				} else {
					Neos[i].Inputs[8] = td / float64(Program.MaxDistanceLook)
					break
				}
			}
			if indexS == Program.MaxDistanceLook-1 {
				Neos[i].Inputs[8] = -1.0
			}

		}

		// Blockage backwards
		// akin to raytracing, which I hate
		// simple populate world array with items them step the direction
		for indexS := 1; indexS < Program.MaxDistanceLook; indexS++ { // going to look to far at anlges but will fix at distance check

			db := Neos[i].Direction
			if db < 180 {
				db += 180
			} else {
				db -= 180
			}

			p := DirectionToStep(db)

			var td float64 // td : temporary distance to a target
			tx := Neos[i].LocationX + (indexS * p.X)
			ty := Neos[i].LocationY + (indexS * p.Y)

			// Check end of world
			if tx >= Program.WorldX || ty >= Program.WorldY || tx < 0 || ty < 0 {
				td = GetDistance(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(tx), float64(ty))
				if td > float64(Program.MaxDistanceLook) {
					Neos[i].Inputs[9] = -1.0
					break
				}
				Neos[i].Inputs[9] = td
				break
			}

			if World[XYtoIndex(tx, ty)] != 0 {
				td = GetDistance(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(tx), float64(ty))
				if td > float64(Program.MaxDistanceLook) {
					Neos[i].Inputs[9] = -1.0
					break
				} else {
					Neos[i].Inputs[9] = td / float64(Program.MaxDistanceLook)
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
			d1 := GetDistance(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(k.X), float64(k.Y))
			a1 := GetAngle(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(k.X), float64(k.Y))
			if d1 < d {
				d = d1
				a = a1
			}
		}
		if d <= float64(Program.MaxDistanceLook) {
			Neos[i].Inputs[5] = d / float64(Program.MaxDistanceLook)
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
			d1 := GetDistance(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(k.LocationX), float64(k.LocationY))
			a1 := GetAngle(float64(Neos[i].LocationX), float64(Neos[i].LocationY), float64(k.LocationX), float64(k.LocationY))
			if d1 < d {
				d = d1
				a = a1
			}
		}
		if d <= float64(Program.MaxDistanceLook) {
			Neos[i].Inputs[6] = d / float64(Program.MaxDistanceLook)
			Neos[i].Inputs[2] = a / 360.0
		} else {
			Neos[i].Inputs[6] = -1.0
			Neos[i].Inputs[2] = -1.0
		}

	}
}

// Step1 : Move the Neo's inputs to the neurons
func Step1() {
	for i := range Neos {
		if i == 0 { // Skip 0
			continue
		}
		for j, n := range Neos[i].Neurons {
			if n.SourceLayer == 0 {
				Neos[i].Neurons[j].InValue = Neos[i].Inputs[n.Source]
			}
		}
	}
}

// Step2 : Propigate out to in and sum and pass through tanh.
func Step2(i int, wg *sync.WaitGroup) {

	buffers <- true

	if i > Program.NumberOfNeos || i < 1 {
		log.Fatalf("Step2 id '%d' is out of bounds", i)
	}

	defer wg.Done()

	//	for i := range Neos { // Loop throught the Neos
	//if i == 0 { // skip 0
	//continue
	//}

	for j := 0; j < Program.NumberOfLayers-1; j++ { // Layer by Layer

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
					Neos[i].Neurons[k].InValue = math.Tanh(m.InValue)
				}
			}
		}
	}

	<-buffers
}

func Step3() {

	for i := 1; i < Program.NumberOfNeos; i++ { // skip 0
		// Check if Neo died from hunger
		if Neos[i].Hunger == 0 {
			Neos[i].Dead = true
			continue
		}
		if Neos[i].Dead {
			continue
		}

		// Moves
	}

}

func probability(p float64) bool {
	if randFloat() <= math.Abs(p) {
		return true
	}
	return false
}
