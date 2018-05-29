package commonerror

type code interface {
	Code() int
}

// Code gets a code from a thing.
//
// Returns 0 or a supplied default code if the thing does not
// implement the code interface.
func Code(thing interface{}, defaultCode ...int) int {
	if c, ok := thing.(code); ok {
		return c.Code()
	}
	if len(defaultCode) > 0 {
		return defaultCode[0]
	}
	return 0
}
