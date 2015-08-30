package main

import(
	"fmt"
	"log"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request){
	http.ServeFile(writer, request, request.URL.Path[1:])
}

func loginHandler(writer http.ResponseWriter, request *http.Request){
	request.ParseForm()
	userEmail := request.Form["email"][0]
	password := request.Form["password"][0]
	if password == "pass"{
		// add session id to database
		fmt.Fprintf(writer, "sessionID")
		log.Println("User: " userEmail + " logged in")
	} 
	
	// get user name and query hashed password will need to hash password server side
}

func main (){
	
	http.HandleFunc("/", indexHandler)	
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Listening on 8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "./ssl/cert.pem", "./ssl/key.pem", nil))
}