package fireworks

type DataDeserialiser[T any] interface {
	Deserialise(data []byte) T
}
