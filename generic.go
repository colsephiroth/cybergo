package cybergo

type GenericResponse[T any] struct {
	Value    []T    `json:"value,omitempty"`
	Count    int    `json:"count,omitempty"`
	NextLink string `json:"nextLink,omitempty"`
}

func Intersection[T comparable](s1, s2 []T) []T {
	var intersection []T
	m := make(map[T]any)
	for _, v := range s1 {
		m[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func Filter[T comparable](s []T, f func(T) bool) []T {
	var filtered []T
	for _, v := range s {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func InSlice[T comparable](item T, slice []T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
