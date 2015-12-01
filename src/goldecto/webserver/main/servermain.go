package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"goldecto"
	"goldecto/webserver"
	"goldecto/webserver/views"
)

func main() {
	conf := goldecto.ConfigFromFile(os.Args[1])
	goldecto.SetResourceDir(os.Args[2])

	webserver.InitRouter()
	
	webserver.AddRoute("samplepage", "GET", "/hello", views.HelloPage)

	listen := fmt.Sprintf("%s:%d", conf.Web.Host, conf.Web.Port)

	fmt.Printf("Listening to %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, webserver.Router))

}
