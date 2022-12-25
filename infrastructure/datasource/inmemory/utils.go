package inmemory

func filter[T comparable](slice []T, filter func(T) bool) (new []T) {
	for _, s := range slice {
		if filter(s) {
			new = append(new, s)
		}
	}
	return
}
