package main

import (
	"fmt"
)

func placeneo() {

	for id := 1; id < Program.NumberOfNeos; id++ {

		p := Point{}
		index := 0
		for {

			p.X = randInt(128)
			p.Y = randInt(128)

			index = XYtoIndex(p.X, p.Y)

			if World[index] == 0 {
				break
			}

		}

		p.FX = float64(p.X)
		p.FY = float64(p.Y)
		Neos[id].Location = p
		World[index] = id
	}
}

func runsim() error {

	for generation = 0; generation < Program.NumberOfGenerations; generation++ {
		fmt.Println("Generation:", generation)

		placeneo()

		for step = 0; step < Program.NumberOfSteps; step++ {
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward0(id)
			}
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward1(id)
			}
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward2(id)
			}
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward3(id)
			}
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward4(id)
			}
			/*

				p0 := parallel.NewParallel()
				for id := 1; id < Program.NumberOfNeos; id++ {
					p0.Register(forward0, id).SetReceivers()
				}
				p0.Run()

				p1 := parallel.NewParallel()
				for id := 1; id < Program.NumberOfNeos; id++ {
					p1.Register(forward1, id).SetReceivers()
				}
				p1.Run()

				p2 := parallel.NewParallel()
				for id := 1; id < Program.NumberOfNeos; id++ {
					p2.Register(forward2, id).SetReceivers()
				}
				p2.Run()

				p3 := parallel.NewParallel()
				for id := 1; id < Program.NumberOfNeos; id++ {
					p3.Register(forward3, id).SetReceivers()
				}
				p3.Run()

				p4 := parallel.NewParallel()
				for id := 1; id < Program.NumberOfNeos; id++ {
					p3.Register(forward4, id).SetReceivers()
				}
				p4.Run()

			*/
			createpng(step)

		}

		for id := 1; id < Program.NumberOfNeos; id++ {
			printgenes(id)
			printneuron(id)
			printio(id)
		}
	}

	return nil
}
