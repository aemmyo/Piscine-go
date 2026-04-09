package lasagna


    const  OvenTime = 40

// TODO: define the 'OvenTime' constant

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
     const  OvenTime = 40
    k := actualMinutesInOven
    s := OvenTime - k
    return s
	//panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    n := numberOfLayers
    x := n * 2
    return x
	//panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
   x := numberOfLayers * 2 +  actualMinutesInOven
    return x
    
	panic("ElapsedTime not implemented")
}
