package proxy

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v8"
)

const backendURL = "http://localhost:8081"
var ctxt = context.Background()
func serialize(h string , header http.Header) string {
   var resp string
   for tit,val:=  range header{
      resp+=tit+" : "+val[0]+"\n"
      log.Println(tit+" : "+val[0])
   }
   response:=h+"\n"+resp
   return response
}
func deserialize(h string) (string,http.Header){ 
   lines := strings.Split(h, "\n")
   if len(lines) == 0 {
      return "", nil
   }

   body := lines[0] 
   headers := http.Header{}
   
   for _, line := range lines[1:] {
      if line == "" { 
         continue
      }
      parts := strings.SplitN(line, ":", 2)
      if len(parts) == 2 {
         key := parts[0]
         value := parts[1]
         headers.Add(key, value)
      }
   }

   return body, headers
}

func HandleRequest(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
       

		key := r.URL.Path
		value, err := rdb.Get(ctxt, key).Result()
		
		if err == redis.Nil {
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

			
         h,err:=io.ReadAll(resp.Body);
         if err!=nil{
            http.Error(w, "could not read response", http.StatusInternalServerError)
            return
         }
          
        
			if err := rdb.Set(ctxt, key, h, 0).Err(); err != nil {
				http.Error(w, "could not set in redis", http.StatusInternalServerError)
				return
			}

			
			w.WriteHeader(resp.StatusCode)
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				http.Error(w, "could not write response", http.StatusInternalServerError)
				return
			}
			return
		} else if err != nil {
			http.Error(w, "could not get from redis", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(value))
	}
}