package syncutil

// runs an async error-returning task in a goroutine and eventually returns the result in a channel
func Async(fn func() error) <-chan error {
	result := make(chan error, 1)
	go func() {
		result <- fn()
		close(result)
	}()
	return result
}
