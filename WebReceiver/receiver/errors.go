package receiver

import "errors"

var (
	ErrorRecord   = errors.New("error record")
	ErrorDataLine = errors.New("error parsing float")
)
