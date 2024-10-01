package proxy

import (
	"io"
	"log"
	"net/http"
)


const backendURL = "http://localhost:8081"
func Handlerequest (w http.ResponseWriter, r*http.Request){
   req,err:=http.NewRequest(
	r.Method,backendURL+r.URL.Path,r.Body,
   )
   if err!=nil{
	http.Error(w,"could not create request",http.StatusInternalServerError)
	return
   }
   req.Header=r.Header
   resp,err:=http.DefaultClient.Do(req)
   if err!=nil{
	http.Error(w,"could not send request",http.StatusInternalServerError)
	return
   }
   defer resp.Body.Close()
   log.Printf("status code: %d", resp.StatusCode)
   w.WriteHeader(resp.StatusCode)
   _,err= io.Copy(w,resp.Body)
   if err!=nil{
	log.Printf("error writing response: %v", err)
	http.Error(w,"could not write response",http.StatusInternalServerError)
	return
   }  
   log.Printf("request: %s", r.URL.Path) 
   
   }	
	
