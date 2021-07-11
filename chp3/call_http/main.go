package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func get_google_basic() (resp *http.Response) {
	r1, _ := http.Get("http://www.google.com/robots.txt")
	return r1
}

func get_andrew_willette_dot_com() (resp *http.Response) {
	r1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	return r1
}

func test_form_submission() (resp *http.Response) {
	form := url.Values{}
	form.Add("foo", "bar")
	r3, err := http.PostForm("https://www.google.com/robots.txt", form)
	if err != nil {
		log.Fatal(err)
	}
	return r3
}

func do_request(http_method string, url string) (resp *http.Response) {
	req, err := http.NewRequest(http_method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	var client http.Client
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("error making request %s\n", err)
	}
	return resp
}

func main() {
	get_google_basic()
}
