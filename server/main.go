package main

import (
	"server/api"
	"server/model/db"
)

func main() {
	db := db.XormConnect()
	defer db.Close()

	finish := make(chan bool)

	go api.ListenAndServe("49200")

	<-finish
}
