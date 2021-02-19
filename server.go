package main
import (
"github.com/gorilla/mux"
"net/http"
"fmt"
"log"	

)

func main(){
 router := mux.NewRouter();

 const port string = ":8000"

 router.HandleFunc("/",func(response http.ResponseWriter, request *http.Request){
	fmt.Fprintln(response,"Up and running...")

 })
 router.HandleFunc("/posts",getPosts).Methods("Get")
 router.HandleFunc("/posts",addPost).Methods("POST")

 log.Println("Server listening on port", port)
// http.ListenAndServe(port, router)
log.Fatalln(http.ListenAndServe(port,router))

}