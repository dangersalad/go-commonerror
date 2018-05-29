package errorinterfaces

type err interface {
	Err() error
}

// Err gets an error from a thing.
//
// Returns thing if the thing does not
// implement the err interface but implements
// the error interface. Returns nil otherwise.
func Err(thing interface{}) error {
	if c, ok := thing.(err); ok {
		return c.Err()
	}
	if e, ok := thing.(error); ok {
		return e
	}
	return nil
}
