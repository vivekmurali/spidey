package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func worker(URL string) {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	//INSERT DATA AND META DATA INTO THE DATABASE FOR THIS PARTICULAR URL
	//
	//

	//PARSE LINKS AND INSERT THEM INTO THE DATABASE

}
