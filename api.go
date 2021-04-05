package boom

import "net/http"

// RenderBadRequest responds with a 400 Bad Request error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderBadRequest(w http.ResponseWriter, message ...interface{}) {
	Render(w, BadRequest(message...))
}

// BadRequest creates the error for a bad request
func BadRequest(message ...interface{}) Err {
	return Boomify(http.StatusBadRequest, message...)
}

// RenderUnauthorized responds with a 401 Unauthorized error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderUnauthorized(w http.ResponseWriter, message ...interface{}) {
	Render(w, Unauthorized(message...))
}

// Unauthorized creates the error for an unauthorized request
func Unauthorized(message ...interface{}) Err {
	return Boomify(http.StatusUnauthorized, message...)
}

// RenderPaymentRequired responds with a 402 Payment Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderPaymentRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, PaymentRequired(message...))
}

// PaymentRequired creates the error for a payment required error
func PaymentRequired(message ...interface{}) Err {
	return Boomify(http.StatusPaymentRequired, message...)
}

// RenderForbidden responds with a 403 Forbidden error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderForbidden(w http.ResponseWriter, message ...interface{}) {
	Render(w, Forbidden(message...))
}

// Forbidden creates the error for a forbidden request
func Forbidden(message ...interface{}) Err {
	return Boomify(http.StatusForbidden, message...)
}

// RenderNotFound responds with a 404 Not Found error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderNotFound(w http.ResponseWriter, message ...interface{}) {
	Render(w, NotFound(message...))
}

// NotFound creates the error for a not found error
func NotFound(message ...interface{}) Err {
	return Boomify(http.StatusNotFound, message...)
}

// RenderMethodNotAllowed responds with a 405 Method Not Allowed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderMethodNotAllowed(w http.ResponseWriter, message ...interface{}) {
	Render(w, MethodNotAllowed(message...))
}

// MethodNotAllowed creates the error for a request that is not allowed
func MethodNotAllowed(message ...interface{}) Err {
	return Boomify(http.StatusMethodNotAllowed, message...)
}

// RenderNotAcceptable responds with a 406 Not Acceptable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderNotAcceptable(w http.ResponseWriter, message ...interface{}) {
	Render(w, NotAcceptable(message...))
}

// NotAcceptable creates the error for a request that is not accepted
func NotAcceptable(message ...interface{}) Err {
	return Boomify(http.StatusNotAcceptable, message...)
}

// RenderProxyAuthRequired responds with a 407 Proxy Authentication Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderProxyAuthRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, ProxyAuthRequired(message...))
}

// ProxyAuthRequired creates the error for a request that requires a proxy auth
func ProxyAuthRequired(message ...interface{}) Err {
	return Boomify(http.StatusProxyAuthRequired, message...)
}

// RenderRequestTimeout responds with a 408 Request Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderRequestTimeout(w http.ResponseWriter, message ...interface{}) {
	Render(w, RequestTimeout(message...))
}

// RequestTimeout creates the error for a request that has timed out
func RequestTimeout(message ...interface{}) Err {
	return Boomify(http.StatusRequestTimeout, message...)
}

// RenderConflict responds with a 409 Conflict error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderConflict(w http.ResponseWriter, message ...interface{}) {
	Render(w, Conflict(message...))
}

// Conflict creates the error for a request that results in a conflict
func Conflict(message ...interface{}) Err {
	return Boomify(http.StatusConflict, message...)
}

// RenderGone responds with a 410 Gone error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderGone(w http.ResponseWriter, message ...interface{}) {
	Render(w, Gone(message...))
}

// Gone creates the error for a resource that is gone
func Gone(message ...interface{}) Err {
	return Boomify(http.StatusGone, message...)
}

// RenderLengthRequired responds with a 411 Length Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderLengthRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, LengthRequired(message...))
}

// LengthRequired creates the error when a length is required
func LengthRequired(message ...interface{}) Err {
	return Boomify(http.StatusLengthRequired, message...)
}

// RenderPreconditionFailed responds with a 412 Precondition Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderPreconditionFailed(w http.ResponseWriter, message ...interface{}) {
	Render(w, PreconditionFailed(message...))
}

// PreconditionFailed creates the error when a precondition fails
func PreconditionFailed(message ...interface{}) Err {
	return Boomify(http.StatusPreconditionFailed, message...)
}

// RenderEntityTooLarge responds with a 413 Request Entity Too Large error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderEntityTooLarge(w http.ResponseWriter, message ...interface{}) {
	Render(w, EntityTooLarge(message...))
}

// EntityTooLarge creates the error when an entity is too large
func EntityTooLarge(message ...interface{}) Err {
	return Boomify(http.StatusRequestEntityTooLarge, message...)
}

// RenderURITooLong responds with a 414 Request-URI Too Large error
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderURITooLong(w http.ResponseWriter, message ...interface{}) {
	Render(w, URITooLong(message...))
}

// URITooLong creates the error when the uri is too long
func URITooLong(message ...interface{}) Err {
	return Boomify(http.StatusRequestURITooLong, message...)
}

// RenderUnsupportedMediaType responds with a 415 Unsupported Media Type error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderUnsupportedMediaType(w http.ResponseWriter, message ...interface{}) {
	Render(w, UnsupportedMediaType(message...))
}

// UnsupportedMediaType creates the error when an unsupported media type is encountered
func UnsupportedMediaType(message ...interface{}) Err {
	return Boomify(http.StatusUnsupportedMediaType, message...)
}

// RenderRangeNotSatisfiable responds with a 416 Requested Range Not Satisfiable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderRangeNotSatisfiable(w http.ResponseWriter, message ...interface{}) {
	Render(w, RangeNotSatisfiable(message...))
}

// RangeNotSatisfiable creates the error when a range is not satisfiable
func RangeNotSatisfiable(message ...interface{}) Err {
	return Boomify(http.StatusRequestedRangeNotSatisfiable, message...)
}

// RenderExpectationFailed responds with a 417 Expectation Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderExpectationFailed(w http.ResponseWriter, message ...interface{}) {
	Render(w, ExpectationFailed(message...))
}

// ExpectationFailed creates the error when an expectation failed
func ExpectationFailed(message ...interface{}) Err {
	return Boomify(http.StatusExpectationFailed, message...)
}

// RenderTeapot responds with a 418 I'm a Teapot error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderTeapot(w http.ResponseWriter, message ...interface{}) {
	Render(w, Teapot(message...))
}

// Teapot creates the error when a teapot error occurs
func Teapot(message ...interface{}) Err {
	return Boomify(http.StatusTeapot, message...)
}

// RenderUnprocessableEntity responds with a 422 Unprocessable Entity error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderUnprocessableEntity(w http.ResponseWriter, message ...interface{}) {
	Render(w, UnprocessableEntity(message...))
}

// UnprocessableEntity creates an unprocessable entity error
func UnprocessableEntity(message ...interface{}) Err {
	return Boomify(http.StatusUnprocessableEntity, message...)
}

// RenderLocked responds with a 423 Locked error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderLocked(w http.ResponseWriter, message ...interface{}) {
	Render(w, Locked(message...))
}

// Locked creates a locked error
func Locked(message ...interface{}) Err {
	return Boomify(http.StatusLocked, message...)
}

// RenderPreconditionRequired responds with a 428 Precondition Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderPreconditionRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, PreconditionRequired(message...))
}

// PreconditionRequired creates an error when a precondition is required
func PreconditionRequired(message ...interface{}) Err {
	return Boomify(http.StatusPreconditionRequired, message...)
}

// RenderTooManyRequests responds with a 429 Too Many Requests error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderTooManyRequests(w http.ResponseWriter, message ...interface{}) {
	Render(w, TooManyRequests(message...))
}

// TooManyRequests creates an error when too many requests are encountered
func TooManyRequests(message ...interface{}) Err {
	return Boomify(http.StatusTooManyRequests, message...)
}

// RenderHeaderFieldsTooLarge responds with a 431 request header fields too large error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderHeaderFieldsTooLarge(w http.ResponseWriter, message ...interface{}) {
	Render(w, HeaderFieldsTooLarge(message...))
}

// HeaderFieldsTooLarge creates an error when header fields are too large
func HeaderFieldsTooLarge(message ...interface{}) Err {
	return Boomify(http.StatusRequestHeaderFieldsTooLarge, message...)
}

// RenderIllegal responds with a 451 Unavailable For Legal Reasons error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderIllegal(w http.ResponseWriter, message ...interface{}) {
	Render(w, Illegal(message...))
}

// Illegal creates an error when a error is encountered from a legal perspective
func Illegal(message ...interface{}) Err {
	return Boomify(http.StatusUnavailableForLegalReasons, message...)
}

// RenderInternal responds with a 500 Internal Server Error error. Alias for boom.Internal.
// Takes an optional message of either type string or type error
// but, will always return a generic message in the response body.
// This behaviour protects the developer from accidentally returning
// sensitive data in the response during a panic.
func RenderInternal(w http.ResponseWriter, message ...interface{}) {
	Render(w, Internal())
}

// Internal creates an internal error
func Internal(message ...interface{}) Err {
	return Boomify(http.StatusInternalServerError)
}

// RenderBadImplementation responds with an internal error
func RenderBadImplementation(w http.ResponseWriter, message ...interface{}) {
	RenderInternal(w)
}

// BadImplementation creates an error to be used when a bad implementation is encountered
func BadImplementation(message ...interface{}) Err {
	return Internal(message...)
}

// RenderNotImplemented responds with a 501 Not Implemented error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderNotImplemented(w http.ResponseWriter, message ...interface{}) {
	Render(w, NotImplemented(message...))
}

// NotImplemented creates an error to be used when an edpoint is not implemented
func NotImplemented(message ...interface{}) Err {
	return Boomify(http.StatusNotImplemented, message...)
}

// RenderBadGateway responds with a 502 Bad Gateway error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderBadGateway(w http.ResponseWriter, message ...interface{}) {
	Render(w, BadGateway(message...))
}

// BadGateway creates an error to be used when a bad gateway is encountered
func BadGateway(message ...interface{}) Err {
	return Boomify(http.StatusBadGateway, message...)
}

// RenderServiceUnavailable .eturns a 503 Service Unavailable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderServiceUnavailable(w http.ResponseWriter, message ...interface{}) {
	Render(w, ServiceUnavailable(message...))
}

// ServiceUnavailable creates an error to be used when a service is unavailable
func ServiceUnavailable(message ...interface{}) Err {
	return Boomify(http.StatusServiceUnavailable, message...)
}

// RenderGatewayTimeout responds with a 504 Gateway Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderGatewayTimeout(w http.ResponseWriter, message ...interface{}) {
	Render(w, GatewayTimeout(message...))
}

// GatewayTimeout creates an error to be used when a gateway timeout occurs
func GatewayTimeout(message ...interface{}) Err {
	return Boomify(http.StatusGatewayTimeout, message...)
}
