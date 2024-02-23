package piscine

func PodiumPosition(podium [][]string) [][]string {
	// Reverse the order of the podium slice.
	for i := 0; i < len(podium)/2; i++ {
		podium[i], podium[len(podium)-i-1] = podium[len(podium)-i-1], podium[i]
	}

	// Return the reversed podium slice.
	return podium
}
