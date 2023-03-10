package xerrors

type Class int

const (
	ClassInternal   Class = 500
	ClassBadRequest Class = 400
	ClassNotFound   Class = 404
)

type ClassError struct {
	class int
	inner error
}

func NewClassError(class int, inner error) *ClassError {
	if inner == nil {
		return nil
	}

	return &ClassError{
		class: class,
		inner: inner,
	}
}

func (e *ClassError) Error() string {
	return e.inner.Error()
}

func (e *ClassError) Unwrap() error {
	return e.inner
}
