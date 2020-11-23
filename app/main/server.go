

package main


import (
	"flycode.go/abacusf/app/tools/debug"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)





func main()  {

    router := mux.NewRouter()

    router.HandleFunc("/sendmail",PContactController2).Methods("POST")

    log.Fatal(http.ListenAndServe("localhost:3001", router))
}




func PContactController2(reponse http.ResponseWriter, req *http.Request)  {

        name := req.FormValue("name")
	    email := req.FormValue("email")
			
	    debug.Log("front<PContactController>  ----> "+name)	
		debug.Log("front<PContactController>  ----> "+email)		
				

		reponse.Header().Set("Access-Control-Allow-Origin", "*")
}