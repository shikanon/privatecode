/*
 * Author: shikanon (shikanon@tensorbytes.com)
 * File Created Time: 2020-05-19 2:35:31
 *
 * Project: auth-service
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
	username := r.URL.Query().Get("username")
	passwd := r.URL.Query().Get("passwd")
	content := fmt.Sprintf("username:%v", username)
	if username == "shikanon" && passwd == "123456" {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}
	fmt.Fprintf(w, content)
	fmt.Println(content)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
