package main

import (
	"fmt"
	"net/http"
	"proxyserver/proxy"

	"github.com/go-redis/redis/v8"
)


func main(){

	rdb:=redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
	})
	 http.HandleFunc("/", proxy.HandleRequest(rdb))
	 fmt.Println("starting proxy server")
	 if err:=http.ListenAndServe(":8080", nil) ;err!= nil {

		 fmt.Println(err)
		 panic(err)

	 }

}