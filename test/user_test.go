package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestSignupUser(t *testing.T) {

	requestURL := fmt.Sprintf("http://localhost:%d", 8080)

	// JSON body
	body := []byte(`{
		"email": "ass1@a.com",
		"password": "111"
	}`)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	defer res.Body.Close()
	// assert.Equal(t, , 13, "they should be equal")

}
