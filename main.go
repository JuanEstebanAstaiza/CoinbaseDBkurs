package main

import "log"

func main() {

	err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}

}
