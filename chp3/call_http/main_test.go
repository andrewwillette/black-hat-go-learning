package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
Basic test, show that 200 exists in response string
**/
func Test_Get_Google_Basic(t *testing.T) {
	resp := get_google_basic()

	defer resp.Body.Close()
	read_values, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error reading bytes from response body")
	}
	read_values_string := string(read_values)
	matched, err := regexp.MatchString("Allow", read_values_string)
	if err != nil {
		log.Fatalln("error doing regex on body")
	}
	assert.Equal(t, true, matched, "Allow exists in response body")
	assert.Equal(t, "200 OK", resp.Status, "response returns 200 status")
}

func Test_some_other(t *testing.T) {
	resp := test_form_submission()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var stringified = string(resp_body)
	match, err := regexp.MatchString("404", stringified)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, true, match, "404 expected in post")
}

func Test_do_request(t *testing.T) {
	resp := do_request("POST", "andrewwillette.com")
	fmt.Println(resp)
}
