/*
 * Author: shikanon (shikanon@tensorbytes.com)
 * File Created Time: 2020-05-20 9:04:52
 *
 * Project: server
 * File: main.go
 * Description:
 *
 */

package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is server")
}

func main() {
	http.HandleFunc("/lookup", indexHandler)
	http.ListenAndServe(":8080", nil)
}
