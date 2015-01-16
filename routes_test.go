package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
)

func TestIndexHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	indexHandler(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "WebRTC Client") {
			t.Errorf("oh shoot didn't get the page! %s", p)
		}

	}
}
