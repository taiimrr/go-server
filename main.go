package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)
func formHandler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseFrom err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request success")
	name := r.FormValue("name")
	address := r.FormValue("add")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s", address)


}
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	router := chi.NewRouter()

	router.Handle("/*", http.StripPrefix("/", fileServer))	
	router.Get("/hello", helloHandler)
	router.Post("/form", formHandler)
	
	srv := &http.Server{
		Handler: router,
		Addr: ":8000",
	}

	log.Printf("Server runnin on port %v", 8000)
	

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	




}