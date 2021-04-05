package boom

import (
	"encoding/json"
	"net/http"
)

// Err holds information about the error to be returned to the client
// during a renderBoom function call or elsewhere
type Err struct {
	ErrorType  string             `json:"error"`
	Message    string             `json:"message"`
	StatusCode int                `json:"statusCode"`
	Data       *map[string]string `json:"data,omitempty"`
}

// Error makes compatable with an error
func (err Err) Error() string {
	msg, _ := json.Marshal(err)
	return err.Message + ": " + string(msg)
}

// Boomify creates a boom error starting with its status code
func Boomify(statusCode int, args ...interface{}) Err {
	errorType := http.StatusText(statusCode)
	message := errorType // should be same as errorType by default

	// determine if an error or string arg was passed in
	// set the message accordingly
	var data *map[string]string
	if l := len(args); l != 0 {
		for _, arg := range args {
			switch converted := arg.(type) {
			case string:
				message = converted
			case error:
				message = converted.Error()
			case map[string]string:
				data = &converted
			}
		}
	}
	return Err{
		errorType,
		message,
		statusCode,
		data,
	}
}

// Render renders a boom error
func Render(w http.ResponseWriter, boomed Err) {
	errString, _ := json.Marshal(boomed)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(boomed.StatusCode)
	_, _ = w.Write(errString)
}
