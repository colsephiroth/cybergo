package pvwa

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type GenericResponse[T any] struct {
	Value    []T    `json:"value,omitempty"`
	Count    int    `json:"count,omitempty"`
	NextLink string `json:"nextLink,omitempty"`
}

func genericGetReturnSingle[R any](pvwa *PVWA, path string, query *url.Values) (*R, error) {
	_path := buildPath(path, query)

	response := new(R)

	pvwa.logIfEnabled("GET " + _path)

	res, err := pvwa.Get(_path)
	if err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}
	defer pvwa.logIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&response); err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	return response, nil
}

func genericGetReturnSlice[R any](pvwa *PVWA, path string, query *url.Values) ([]*R, error) {
	_path := buildPath(path, query)

	var response []*R

	for {
		pvwa.logIfEnabled("GET " + _path)

		data := new(GenericResponse[*R])

		res, err := pvwa.Get(_path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		pvwa.logIfError(res.Close)

		response = append(response, data.Value...)

		if data.NextLink != "" {
			_path = data.NextLink
		} else {
			break
		}
	}

	return response, nil
}

func genericPostReturnSingle[T, R any](pvwa *PVWA, path string, query *url.Values, data T) (*R, error) {
	_path := buildPath(path, query)

	response := new(R)

	pvwa.logIfEnabled(fmt.Sprintf("POST %s %#v", _path, data))

	_data, err := json.Marshal(data)
	if err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	res, err := pvwa.Post(_path, _data)
	if err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}
	defer pvwa.logIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&response); err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	return response, nil
}

func genericPatchReturnSingle[T, R any](pvwa *PVWA, path string, query *url.Values, data T) (*R, error) {
	_path := buildPath(path, query)

	response := new(R)

	pvwa.logIfEnabled(fmt.Sprintf("PATCH %s %#v", _path, data))

	_data, err := json.Marshal(data)
	if err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	res, err := pvwa.Patch(_path, _data)
	if err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}
	defer pvwa.logIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&response); err != nil {
		pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	return response, nil
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

type Option[T any] struct {
	value T
	valid bool
}

func (o *Option[T]) Some(t T) Option[T] {
	o.valid = true
	o.value = t
	return *o
}

func (o *Option[T]) Value() T {
	return o.value
}

func (o *Option[T]) Valid() bool {
	return o.valid
}

func httpStatusSuccess(code int) bool {
	return code >= 200 && code < 300
}
