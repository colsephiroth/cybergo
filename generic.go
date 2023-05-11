package cybergo

type GenericResponse[T any] struct {
	Value    []T    `json:"value,omitempty"`
	Count    int    `json:"count,omitempty"`
	NextLink string `json:"nextLink,omitempty"`
}
