package main

import (
	"log"
	"net/http"
)

func main(){
	log.Println("Listening to port : 8082")
	mux := http.NewServeMux()


	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		log.Printf("%v", r.RequestURI)
		w.Write([]byte("pong"))
		
	})
	if err := http.ListenAndServe(":8082", mux); err != nil{
		log.Fatal(err)
	}
}