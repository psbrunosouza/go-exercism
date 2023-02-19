package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdsTotal := 0

	for i := 0; i < len(birdsPerDay); i++ {
		birdsTotal += birdsPerDay[i]
	}

	return birdsTotal
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsTotal := 0

	for i := 0; i < len(birdsPerDay); i++ {
		if (i >= 0 && i <= 6) && week == 1 {
			birdsTotal += birdsPerDay[i]
		} else if (i >= 7 && i <= 13) && week == 2 {
			birdsTotal += birdsPerDay[i]
		}
	}

	return birdsTotal
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	birdsPerDayFixed := birdsPerDay

	for i := 1; i <= len(birdsPerDay); i++ {
		if i%2 != 0 {
			birdsPerDayFixed[i-1] += 1
		}
	}

	return birdsPerDayFixed
}
