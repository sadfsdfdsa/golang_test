package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRunMain(t *testing.T) {
	go main()
}

func TestGetRequest(t *testing.T) {
	client := &http.Client{}

	r, _ := http.NewRequest("GET", "http://localhost:8181", nil)

	resp, err := client.Do(r)
	if err != nil {
		t.Errorf("Error with response")
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error with body")
	}
	bodyString := string(bodyBytes)

	if bodyString != "Index Page" {
		t.Errorf("Error with body content")
	}
}

func TestProducts(t *testing.T) {
	// go main()
	client := &http.Client{}

	r, _ := http.NewRequest("GET", "http://localhost:8181/products/1", nil)

	resp, err := client.Do(r)
	if err != nil {
		t.Errorf("Error with response")
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error with body")
	}
	bodyString := string(bodyBytes)

	if bodyString != "id=1" {
		t.Errorf("Error with body content")
	}
}

func TestArticlesGet(t *testing.T) {
	// go main()
	client := &http.Client{}

	r, _ := http.NewRequest("GET", "http://localhost:8181/articles/1", nil)

	resp, err := client.Do(r)
	if err != nil {
		t.Errorf("Error with response")
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error with body")
	}
	bodyString := string(bodyBytes)

	if bodyString != "" {
		t.Errorf("Error with body content")
	}
}


func TestArticlesPost(t *testing.T) {
	// go main()
	client := &http.Client{}

	r, _ := http.NewRequest("POST", "http://localhost:8181/articles/1", nil)

	resp, err := client.Do(r)
	if err != nil {
		t.Errorf("Error with response")
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error with body")
	}
	bodyString := string(bodyBytes)

	if bodyString != "id=1" {
		t.Errorf("Error with body content")
	}
}
