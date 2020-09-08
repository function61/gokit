package httputils

import (
	"fmt"
	"net/http"
)

// when you don't want to break out gorilla/mux just to get support for routes per method..
type MethodMux struct {
	GET    *http.ServeMux
	HEAD   *http.ServeMux
	POST   *http.ServeMux
	PUT    *http.ServeMux
	DELETE *http.ServeMux
}

// interface assertion
var _ = (http.Handler)(&MethodMux{})

func NewMethodMux() *MethodMux {
	return &MethodMux{
		GET:    http.NewServeMux(),
		HEAD:   http.NewServeMux(),
		POST:   http.NewServeMux(),
		PUT:    http.NewServeMux(),
		DELETE: http.NewServeMux(),
	}
}

func (m *MethodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		m.GET.ServeHTTP(w, r)
	case http.MethodHead:
		m.HEAD.ServeHTTP(w, r)
	case http.MethodPost:
		m.POST.ServeHTTP(w, r)
	case http.MethodPut:
		m.PUT.ServeHTTP(w, r)
	case http.MethodDelete:
		m.DELETE.ServeHTTP(w, r)
	default:
		http.Error(w, fmt.Sprintf("unsupported method: %s", r.Method), http.StatusBadRequest)
	}
}
