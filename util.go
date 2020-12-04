package main

import "log"

func checkErr(err error) {
	if err != nil {
		log.Fatalf("ERROR: %s",err)
	}
}
