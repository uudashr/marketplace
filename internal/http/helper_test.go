package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func httpGet(h http.Handler, path string) *http.Response {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Result()
}

func httpPost(h http.Handler, path string, body interface{}) *http.Response {
	bodyReader, err := makeReader(body)
	if err != nil {
		panic(err)
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, path, bodyReader)
	if bodyReader != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	h.ServeHTTP(rec, req)
	return rec.Result()
}

// convert body to json as reader
func makeReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	switch v := body.(type) {
	case string:
		// return as is
		return strings.NewReader(v), nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		return bytes.NewReader(b), nil
	}
}
