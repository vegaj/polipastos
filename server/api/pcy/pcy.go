package pcy

import (
	"encoding/json"
	"net/http"
)

//HTTPPcy is a policy that describes an scenario to ve validated on the request
type HTTPPcy interface {
	Validate(r *http.Request) error
}

//Handler will wrap a http.HandleFunc and will execute the handler if
//all the policies specified are satisfied.
//The policies will be executed in order of addition
type Handler struct {
	law            []HTTPPcy
	successHandler http.HandlerFunc
	errorHandler   http.HandlerFunc
}

//NewHandler creates a new Handler with some given policies
func NewHandler(onSuccess, onError http.HandlerFunc, law ...HTTPPcy) Handler {
	return Handler{
		law:            law,
		successHandler: onSuccess,
		errorHandler:   onError,
	}
}

//ServeHTTP satisfies the interface http.HandlerFunc
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var err error

	for _, p := range h.law {
		if err = p.Validate(r); err != nil {
			break
		}
	}

	if err == nil {
		h.successHandler(w, r)
	} else {
		h.errorHandler(w, r)
	}
}

var (
	mPolicyViolation = map[string]string{"error": "policy violation"}
)

//DefaultJSONErrorHandler will return an application/json encoded response with a "policy violation" error
func DefaultJSONErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mPolicyViolation)
}
