package convert

func Change(Mainstring string) string { // In order to exclude the return carriage from the thinkertoy.txt
	var sentence string
	for i := 0; i < len(Mainstring); i++ {
		if Mainstring[i] != '\r' {
			sentence += string(Mainstring[i])
		}
	}
	return sentence
}
