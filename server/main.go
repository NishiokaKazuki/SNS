package main

import (
	"fmt"
	"net/http"
)

func jsonResponse(rw http.ResponseWriter, req *http.Request) {
	response := []byte(`
        {
          "status": "success",
          "user":
            {
              "id": "0903",
              "name": "kazupoyo",
              "birthday": "1998-09-03"
            }
        }
    `)

	defer func() {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(response))
	}()
}

func main() {
	http.HandleFunc("/json", jsonResponse)
	http.ListenAndServe(":8080", nil)
}
