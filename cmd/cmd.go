package cmd

func Execute[T comparable](flagValue T, f func(t T) error) error {
	return f(flagValue)
}
