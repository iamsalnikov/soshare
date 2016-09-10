package vk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOKResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "VK.Share.count(1, 37);")
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://etokavkaz.ru")
	if err != nil {
		t.Error("http://etokavkaz.ru must be shared")
	}

	if result != 37 {
		t.Error("Expected 37, got ", result, err)
	}
}

func TestErrorResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "")
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://etokavkaz.ru")
	if err == nil {
		t.Error("http://etokavkaz.ru must not be shared")
	}

	if result > -1 {
		t.Error("Result must be -1")
	}
}
