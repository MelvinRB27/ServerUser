package handlers

import(
	"net/http"
)

//routes for the serve
func Route (mux *http.ServeMux){
	var handler Handlers
	h := NewData(handler)

	//ROUTES OF THE ENDPOINTS
	mux.HandleFunc("/v1/login", h.login)
	mux.HandleFunc("/v1/register", h.register)
	mux.HandleFunc("/v1/update", h.update)
}