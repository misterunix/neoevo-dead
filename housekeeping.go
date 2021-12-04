package main

// clearworld : Set all locations in the world slices to 0
func clearworld() {
	for i := 0; i < Program.WorldSize; i++ {
		World[i] = 0
		WorldTmp[i] = 0
	}
}
