package errors

type ClientError struct {
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

const (
	Invalid       string = "invalid"
	MissingField  string = "missing_field"
	AlreadyExists string = "already_exists"
)

type ValidationError struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

func (ce ClientError) HasError() bool {
	return len(ce.Message) > 0 || len(ce.Errors) > 0
}

func (ce *ClientError) Error() string {
	return ce.Message
}
