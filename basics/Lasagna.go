package basics

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
