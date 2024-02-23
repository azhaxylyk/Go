package piscine

func ReverseMenuIndex(menu []string) []string {
	reversedMenu := make([]string, len(menu))

	for i, j := 0, len(menu)-1; i < len(menu); i, j = i+1, j-1 {
		reversedMenu[i] = menu[j]
	}

	return reversedMenu
}
