package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
     var y int = productionRate 
     k := float64(y)
    return k * ( successRate/100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var y int = productionRate 
     k := float64(y)
    z := k * ( successRate/100)
    
   return int(z/60)
}

func CalculateCost(carsCount int) uint {
    v := carsCount / 10
    t := carsCount % 10
    j := (v * 95000) 
    k := (t * 10000)
    return uint(j+k)
}
