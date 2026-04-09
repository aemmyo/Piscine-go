package lasagna


    const  OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
     const  OvenTime = 40
    k := actualMinutesInOven
    s := OvenTime - k
    return s
}
func PreparationTime(numberOfLayers int) int {
    n := numberOfLayers
    x := n * 2
    return x
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
   x := numberOfLayers * 2 +  actualMinutesInOven
    return x
}
