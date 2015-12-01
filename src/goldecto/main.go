package main

import (
	"fmt"
	"os"
	
	"goldecto/util"
	"goldecto/webserver"
)

func main() {
	conf := util.ConfigFromFile(os.Args[1])
	fmt.Println(conf)

	util.SetResourceDir(os.Args[2])

	webserver.StartServer(conf)
}
