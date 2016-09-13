package googleplus

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
		fmt.Fprintln(w, `{"id": "p", "result": {"kind": "pos#plusones", "id": "http://twitter.com/", "isSetByViewer": false, "metadata": {"type": "URL","globalCounts": {"count": 71446}},"abtk": "AEIZW7QsVzDXx9uSUH2bLNB0302eMeBhG19VV3lsgjtxMf48qJBRbICpimKZknwq8TkspTNlWPTk"}}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://twitter.com")
	if err != nil {
		t.Error("twitter.com must be shared. Got", err)
	}

	if result != 71446 {
		t.Error("Expected 71446, got", result)
	}
}

func TestBadResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"error": {"code": 400, "message": "Invalid Value", "data": [{"domain": "global", "reason": "invalid",	"message": "Invalid Value"}]}, "id": "p"}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://")
	if err == nil {
		t.Error("twitter.com must not be shared. Got", err)
	}

	if result > -1 {
		t.Error("Expected -1, got", result)
	}
}

func TestBadStatusResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"id": "p", "result": {"kind": "pos#plusones", "id": "http://twitter.com/", "isSetByViewer": false, "metadata": {"type": "URL","globalCounts": {"count": 71446}},"abtk": "AEIZW7QsVzDXx9uSUH2bLNB0302eMeBhG19VV3lsgjtxMf48qJBRbICpimKZknwq8TkspTNlWPTk"}}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.GetShareCount("http://twitter.com")
	if err == nil {
		t.Error("twitter.com must not be shared. Got", err)
	}

	if result > -1 {
		t.Error("Expected -1, got", result)
	}
}
