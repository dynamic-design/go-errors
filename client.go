package errors

/**
 * A ClientError is returned to the requesting client if client passed
 * invalid data. It should return a humanreadable explanation (Message)
 * and possibly a list of validation errors.
 *
 * It should be paired with a statuscode 4xx.
 */
type ClientError struct {
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

func (ce ClientError) HasError() bool {
	return len(ce.Message) > 0 || len(ce.Errors) > 0
}

func (ce *ClientError) Error() string {
	return ce.Message
}
