package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rwinkhart/go-boilerplate/back"
)

// GenericResponse defines a generic status and message JSON response.
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// WriteHTTPError writes an error status and message to an http ResponseWriter.
func WriteHTTPError(w *http.ResponseWriter, message string, responseCode int) {
	jsonError := GenericResponse{
		Status:  "error",
		Message: message,
	}

	// ERROR IGNORED
	// Reason: Used in error handling path, thus this needs to be unable
	// to fail (there would be no way to handle it gracefully).
	output, _ := json.Marshal(jsonError)

	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Write(output)
}

// WriteHTTPResponse writes a custom response to an http ResponseWriter.
func WriteHTTPResponse(w *http.ResponseWriter, unmarshaledResponse any) {
	output, err := json.Marshal(unmarshaledResponse)
	if err != nil {
		WriteHTTPError(w, "Failed to marshal JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(output)
}

// LogAndVerifyRequest logs incoming HTTP requests
// and ensures the methods match what is expected.
// Returns: badMethod (bool) - if true, the method is not allowed and the handler should return.
func LogAndVerifyRequest(w *http.ResponseWriter, r *http.Request, expectedMethod string) bool {
	if r.Method != expectedMethod {
		WriteHTTPError(w, "Method not allowed", http.StatusMethodNotAllowed)
		PrintLog(fmt.Sprintf("BAD method (expected: %s, got: %s) for request received on %s from %s", expectedMethod, r.Method, r.URL.Path, r.RemoteAddr), back.AnsiError)
		return true
	}
	PrintLog(fmt.Sprintf("Request received on %s from %s", r.URL.Path, r.RemoteAddr), "")
	return false
}
