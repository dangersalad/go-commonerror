package errorinterfaces

type timeout interface {
	Timeout() bool
}

// Timeout gets if thing is a timeout.
//
// Returns false if the thing does not
// implement the timeout interface.
func Timeout(thing interface{}) bool {
	if c, ok := thing.(timeout); ok {
		return c.Timeout()
	}
	return false
}
