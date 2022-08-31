package fireworks

type ValueProvider interface {
	Get() (string, bool)
}
