package controller
import (
	"encoding/json"
	"net/http"
	"go-APIs/views" // 👈 This matches your module name in go.mod
)
func ping() http.HandlerFunc {  
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//call views.Response 
			data := views.Respose{
				Code: http.StatusOK,
				 Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} 
	}
}