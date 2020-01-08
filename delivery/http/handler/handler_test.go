package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAbout tests GET /about request handler
func TestAbout(t *testing.T) {

	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	About(httprr, req)
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "About" {
		t.Errorf("want the body to contain the word %q", "about")
	}
}

func TestContact(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contact", Contact)
	testServ := httptest.NewTLSServer(mux)
	defer testServ.Close()

	testClient := testServ.Client()
	url := testServ.URL

	resp, err := testClient.Get(url + "/cntact")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Contact" {
		t.Errorf("want the body to contain the word %q", "Contact")
	}
}
