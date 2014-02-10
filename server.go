package errors

/**
 * A ServerError is returned to the client if an error arouse
 * that was not caused by the client but from the server itself.
 *
 * It should be paired with a statuscode 5xx.
 */
type ServerError struct {
	Message string `json:"message"`
}
