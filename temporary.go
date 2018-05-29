package errorinterfaces

type temporary interface {
	Temporary() bool
}

// Temporary gets if a thing is temporary.
//
// Returns false if the thing does not
// implement the temporary interface.
func Temporary(thing interface{}) bool {
	if c, ok := thing.(temporary); ok {
		return c.Temporary()
	}
	return false
}
