package fireworks

import "encoding/json"

type JsonDataDeserialiser[T any] struct{}

func (this JsonDataDeserialiser[T]) Deserialise(data []byte) T {
	var t T
	json.Unmarshal(data, &t)

	return t
}
