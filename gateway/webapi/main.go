//	brian taylor vann
//	taylorvann dot com

package main

import (
	"log"

	"webapi/certificatesx"
	"webapi/server"
)

func main() {
	output, errCerts := certificatesx.Create()
	if errCerts != nil {
		log.Println(output)
	}

	server.CreateServer()
	log.Println("created server")
}
