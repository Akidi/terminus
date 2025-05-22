// import path: terminus/internal/common
// file path: ./internal/common/common.go
package common

import "encoding/json"

type JSONConvertible[T any] interface {
	AsJSON() T
	ToJSON() ([]byte, error)
}

func ToJSON[T any](v JSONConvertible[T]) ([]byte, error) {
	return json.Marshal(v.AsJSON())
}