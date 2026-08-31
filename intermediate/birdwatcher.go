package intermediate

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, bird := range birdsPerDay {
		total += bird
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	var start int = 0
	if week > 1 {
		start = (week - 1) * 7
	}
	for i, bird := range birdsPerDay {
		if i >= start && i <= start+7 {
			total += bird
		}
		i++
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i == 0 || i%2 == 1 {
			birdsPerDay[i] += 1
		}
		i++
	}
	return birdsPerDay
}
