package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/", invite)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func invite(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		request_url := "https://slack.com/api/users.admin.invite"
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		form := url.Values{
			"email": {r.FormValue("email")},
			"token": {getToken()},
		}

		body := bytes.NewBufferString(form.Encode())
		resp, err := http.Post(request_url, "application/x-www-form-urlencoded", body)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(resp)
		}

		body_byte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(body_byte))
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func getToken() string {
	val, ok := os.LookupEnv("TOKEN")
	if !ok {
		return "ERR: TOKEN not set."
	} else {
		return val
	}
}
