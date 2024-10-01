package proxy

import (
	"context"
	"io"
	
	"net/http"

	"github.com/go-redis/redis/v8"
)


const backendURL = "http://localhost:8081"
var ctxt = context.Background()
func HandleRequest(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path
		value, err := rdb.Get(ctxt, key).Result()
		if err != nil {
			http.Error(w, "could not get from redis", http.StatusInternalServerError)
			return
		}
		if value != "" {
			w.Write([]byte(value))
			return
		}

		req, err := http.NewRequest(r.Method, backendURL+r.URL.Path, r.Body)
		if err != nil {
			http.Error(w, "could not create request", http.StatusInternalServerError)
			return
		}
		req.Header = r.Header
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "could not send request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		if err := rdb.Set(ctxt, key, resp.Status, 0).Err(); err != nil {
			http.Error(w, "could not set in redis", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(resp.StatusCode)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, "could not write response", http.StatusInternalServerError)
			return
		}
	}

}
	
