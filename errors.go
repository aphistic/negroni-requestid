package requestid

import "errors"

var (
	ErrMissing = errors.New("Missing request ID value in context")
	ErrInvalid = errors.New("Invalid type found for request ID in context")
)
