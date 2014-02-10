package errors

/**
 * A ValidationError should describe which field that caused the error (Field),
 * what type of error it was (Code) and a human readable explanation (Message).
 *
 * Example:
 * [
 * 		Field: "name",
 * 		Code: errors.Invalid,
 * 		Message: "Name is < 5 characters long",
 * ]
 */
type ValidationError struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
