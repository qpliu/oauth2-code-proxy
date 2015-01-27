package app

import (
	"io"
	"net/http"
	"strings"

	"appengine"
	"appengine/urlfetch"
)

func RegisterService(service Service) {
	http.HandleFunc(service.ProxyPath(), func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		service.AddSecret(r.Form)
		req, err := http.NewRequest("POST", service.URL(), strings.NewReader(r.Form.Encode()))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if username, password, ok := service.BasicAuth(); ok {
			req.SetBasicAuth(username, password)
		}
		resp, err := urlfetch.Client(appengine.NewContext(r)).Do(req)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)
		if _, err = io.Copy(w, resp.Body); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})
}

func Register(services ...Service) {
	for _, service := range services {
		RegisterService(service)
	}
}
