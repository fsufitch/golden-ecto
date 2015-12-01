package webserver

import (
	"fmt"
	"log"
	"net/http"

	"goldecto/util"
	"goldecto/webserver/views"
)

func StartServer(conf util.GoldEctoConfig) {
	InitRouter()
	
	AddRoute("samplepage", "GET", "/hello", views.HelloPage)

	listen := fmt.Sprintf("%s:%d", conf.Web.Host, conf.Web.Port)

	fmt.Printf("Listening to %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, Router))

}
