package clientError

type ClientError struct {
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}

type ValidationError struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

func (ce ClientError) HasError() bool {
	return len(ce.Message) > 0 || len(ce.Errors) > 0
}
