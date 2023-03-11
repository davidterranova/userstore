package xerrors

type Class int

const (
	ClassInternal   Class = 500
	ClassBadRequest Class = 400
	ClassNotFound   Class = 404
	ClassConflict   Class = 409
)

type ClassError struct {
	class Class
	inner error
}

func NewClassError(class Class, inner error) *ClassError {
	if inner == nil {
		return nil
	}

	return &ClassError{
		class: class,
		inner: inner,
	}
}

func (e *ClassError) Class() Class {
	return e.class
}

func (e *ClassError) Error() string {
	return e.inner.Error()
}

func (e *ClassError) Unwrap() error {
	return e.inner
}
