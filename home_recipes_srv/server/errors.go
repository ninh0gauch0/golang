package server

import "fmt"

// HRSError interface
type HRSError interface {
	ShowError() string
}

// FunctionalError Error
type FunctionalError struct {
	Err error
}

// ShowError - Shows
func (fe *FunctionalError) ShowError() string {
	return fmt.Sprintf("[ERROR] A functional error occured: %v", fe.Err)
}

// SetError function
func (fe *FunctionalError) SetError(err error) {
	fe.Err = err
}

// GetError function
func (fe *FunctionalError) GetError() error {
	return fe.Err
}
