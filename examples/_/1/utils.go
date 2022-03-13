package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func parseURL(r *http.Request, path string) string {
	return strings.ToLower(strings.Split(r.Proto, "/")[0]) + "://" + r.Host + path
}

func postRequest(url string, contentType string, data interface{}) (*http.Response, error) {
	buf, err := encodeBuffer(data)
	if err != nil {
		return nil, err
	}
	return http.Post(url, contentType, buf)
}

func encodeBuffer(data interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	return &buf, json.NewEncoder(&buf).Encode(data)
}
