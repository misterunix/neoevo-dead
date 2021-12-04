package main

func decode(gene uint32) Neuron {

	neu := Neuron{}
	weightI := (gene & 0x0000FFFF)               // weightI : unsigned integer value of the weight
	weight := (float64(weightI)/65535.0)*8 - 4.0 // weight : neurons weight -4.0 to +4.0

	src := ((gene & 0x1F000000) >> 24) % INPUTCOUNT
	srcl := ((gene & 0xE0000000) >> 29) % uint32(Program.NumberOfLayers)

	dst := (gene & 0x001F0000) >> 16 % OUTPUTCOUNT
	dstl := (gene & 0x00E00000) >> 21 % uint32(Program.NumberOfLayers)

	neu.Weight = weight
	neu.SourceID = int(src)
	neu.SourceLayer = int(srcl)
	neu.OutID = int(dst)
	neu.OutLayer = int(dstl)

	return neu
}
