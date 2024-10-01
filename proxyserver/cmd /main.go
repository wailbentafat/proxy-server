package main

import ("net/http"
		"proxyserver/proxy"
		"fmt"
	)


func main(){

	 http.HandleFunc("/", proxy.Handlerequest)
	 fmt.Println("starting proxy server")
	 if err:=http.ListenAndServe(":8080", nil) ;err!= nil {
		
		 fmt.Println(err)
		 panic(err)

	 }

}