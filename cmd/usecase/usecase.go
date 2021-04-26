package usecase

import "fmt"

type ResourceNotFoundError struct {
	Err error
}

func NewResourceNotFoundError(err error) *ResourceNotFoundError {
	return &ResourceNotFoundError{Err: err}
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprint(e.Err)
}

func (e *ResourceNotFoundError) Unwrap() error {
	return e.Err
}
