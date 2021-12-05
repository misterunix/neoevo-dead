package main

import "github.com/buptmiao/parallel"

func runsim() error {

	for generation = 0; generation < Program.NumberOfGenerations; generation++ {

		for step = 0; step < Program.NumberOfSteps; step++ {

			p0 := parallel.NewParallel()
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward0(id)
			}
			p0.Run()

			p1 := parallel.NewParallel()
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward1(id)
			}
			p1.Run()

			p2 := parallel.NewParallel()
			for id := 1; id < Program.NumberOfNeos; id++ {
				forward2(id)
			}
			p2.Run()

		}

	}

	return nil
}
