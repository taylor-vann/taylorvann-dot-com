//	brian taylor vann
//	briantaylorvann dot com

package main

import (
	"log"

	"webapi/certificatesx"
	"webapi/server"
)

func main() {
	certOutput, errCertOutput := certificatesx.Create()
	if errCertOutput != nil {
		log.Println("error creating certificates")
		log.Println(errCertOutput)
		log.Println(certOutput)
	}

	server.CreateServer()
}
