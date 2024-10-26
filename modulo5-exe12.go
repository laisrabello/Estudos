package main

func Well(x []string) string {
	cont := 0

	for _, idea := range x {
		if idea == "good" {
			cont += 1
		}
	}

	switch {
	case cont == 1 || cont == 2:
		return "Publish!"
	case cont > 2:
		return "I smell a series!"
	default:
		return "Fail!"
	}

}
