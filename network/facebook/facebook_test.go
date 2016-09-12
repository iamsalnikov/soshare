package facebook

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
		fmt.Fprintln(w, `{"og_object": {"id": "1213871481966215", "title": "Alexey Salnikov", "type": "website", "updated_time": "2016-08-23T13:57:50+0000"}, "share": { "comment_count": 0, "share_count": 90}, "id": "http://iamsalnikov.ru"}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://iamsalnikov.ru")
	if err != nil {
		t.Error("http://iamsalnikov.ru must be shared")
	}

	if result != 90 {
		t.Error("Expected 90, got ", result)
	}
}

func TestErrorResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"error": {"message": "(#803) Some of the aliases you requested do not exist: http:", "type": "OAuthException", "code": 803, "fbtrace_id": "GBNmCt3sywl"}}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://iamsalnikov.ru")
	if err == nil {
		t.Error("http: must not be shared")
	}

	if result > -1 {
		t.Error("Expected -1, got ", result)
	}
}
