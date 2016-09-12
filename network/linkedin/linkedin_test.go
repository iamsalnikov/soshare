package linkedin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOKResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"count": 349892, "fCnt": "349K", "fCntPlusOne": "349K", "url": "http://iamsalnikov.ru"}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://iamsalnikov.ru")
	if err != nil {
		t.Error("http://iamsalnikov.ru must be shared")
	}

	if result != 349892 {
		t.Error("Expected 349892, got ", result)
	}
}

func TestErrorResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `Invalid URL parameter: iamsalnikov.ru`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("iamsalnikov.ru")
	if err == nil {
		t.Error("iamsalnikov.ru must not be shared")
	}

	if result > -1 {
		t.Error("Expected -1, got ", result)
	}
}

func TestBadStatusResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"count": 349892, "fCnt": "349K", "fCntPlusOne": "349K", "url": "http://iamsalnikov.ru"}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("iamsalnikov.ru")
	if err == nil {
		t.Error("iamsalnikov.ru must not be shared")
	}

	if result > -1 {
		t.Error("Expected -1, got ", result)
	}
}
