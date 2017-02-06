package requestid

import (
	"context"
	"net/http"

	"github.com/satori/go.uuid"
)

type contextKey int

const (
	requestIDKey contextKey = iota
)

type RequestIDGen func() string

func defaultIDGen() string {
	return uuid.NewV4().String()
}

type Middleware struct {
	GenerateID RequestIDGen
	XHeader    string
}

func NewMiddleware() *Middleware {
	return &Middleware{
		GenerateID: defaultIDGen,
		XHeader:    "X-Request-ID",
	}
}

func (r *Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	reqID := r.GenerateID()
	ctxt := context.WithValue(req.Context(), requestIDKey, reqID)

	req = req.WithContext(ctxt)

	if r.XHeader != "" {
		rw.Header().Set(r.XHeader, reqID)
	}

	next(rw, req)
}

func FromContext(ctxt context.Context) (string, error) {
	if ctxt == nil {
		return "", ErrMissing
	}

	id := ctxt.Value(requestIDKey)
	if id == nil {
		return "", ErrMissing
	}

	if idStr, ok := id.(string); ok {
		return idStr, nil
	}

	return "", ErrInvalid
}
