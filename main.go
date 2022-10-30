package main

import (
	"flagger-backend/routers"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	routers.Run()
}
