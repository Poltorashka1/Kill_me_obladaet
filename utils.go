package Project

func FindPresence(word string, slice []string) bool {
	for _, v := range slice {
		if word == v {
			return true
		}
	}
	return false
}
