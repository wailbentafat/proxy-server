package main

import (
	"fmt"
	"net/http"
	"proxyserver/proxy"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"proxyserver/db"
)


func main(){
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	
	if err := db.AutoMigrate(&models.Server{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}	

	rdb:=redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
	})
	 http.HandleFunc("/", proxy.HandleRequest(rdb))
	 log.Println("starting proxy server")
	 if err:=http.ListenAndServe(":8080", nil) ;err!= nil {

		 fmt.Println(err)
		 panic(err)

	 }

}