package main

import (
	"fmt"
)

// fitness : evaluate if the gene is viable
func fitness(id int, gn int) bool {

	g := Neo[id].Genes[gn]
	//s1 := (g & 0xFF000000 >> 24) & 0x7f
	t1 := ((g & 0xFF000000 >> 24) & 0x80) >> 7
	if t1 == 1 {
		if hiddencount == 0 {
			return false
		}
	}

	//s2 := (g & 0x00FF0000 >> 16) & 0x7f
	t2 := ((g & 0x00FF0000 >> 16) & 0x80) >> 7
	if t2 == 0 {
		if hiddencount == 0 {
			return false
		}
	}

	return true
}

func buildNeurons(id int) error {
	if id >= BaseNeoCount || id < 0 {
		return fmt.Errorf("buildNeurons id '%d' is out of bounds", id)
	}

	for _, g := range Neo[id].Genes {

		//mt.Println(g)

		var src int
		var snk int
		var srcflag bool
		var snkflag bool
		weightI := (g & 0x0000FFFF)
		weight := (float64(weightI)/65535.0)*2 - 1.0

		//fmt.Println(weight, weightI)

		s := (g & 0xFF000000 >> 24) & 0x7f
		t := ((g & 0xFF000000 >> 24) & 0x80) >> 7
		if t == 0 {
			// input
			src = s % inputcount
			srcflag = false
		} else {
			// hidden
			if hiddencount != 0 {
				src = s % hiddencount // + inputcount + outputcount
				srcflag = true
				//src = (s % Neos[id].Hidden) + 40 // hidden start at 40
			} else {
				src = -1
			}
		}

		//fmt.Printf("s:%d t:%v src:%d\n", s, t, src)

		s = (g & 0x00FF0000 >> 16) & 0x7f
		t = ((g & 0x00FF0000 >> 16) & 0x80) >> 7
		if t == 0 {
			// hidden
			if hiddencount != 0 {
				snk = s % hiddencount // + inputcount + outputcount
				snkflag = false
			} else {
				snk = -1
			}
		} else {
			// out
			snk = s % outputcount //+ inputcount
			snkflag = true
		}

		//fmt.Printf("s:%d t:%v snk:%d\n", s, t, snk)

		if snk != -1 && src != -1 {

			tn := Neuron{
				Source:  src,
				Sink:    snk,
				Weight:  weight,
				SrcFlag: srcflag,
				SnkFlag: snkflag,
			}
			Neo[id].Neurons = append(Neo[id].Neurons, tn)
			//fmt.Printf("tn: %+v\n", tn)

		}
	}

	return nil

}
