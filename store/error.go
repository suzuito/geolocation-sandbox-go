package store

// Error ...
type Error interface {
	Raw() error
}

// ErrorImpl ...
type ErrorImpl struct {
	err error
}

// Raw ...
func (e *ErrorImpl) Raw() error {
	return e.err
}

// NewErrorImpl ...
func NewErrorImpl(err error) *ErrorImpl {
	return &ErrorImpl{
		err: err,
	}
}
