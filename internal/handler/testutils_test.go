package handler

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testServer struct {
	*httptest.Server
}

func newTestServer(handler http.Handler) *testServer {
	return &testServer{
		Server: httptest.NewServer(handler),
	}
}

func (s *testServer) getByNumber(t *testing.T, url string) (resp *http.Response, err error) {
	resp, err = http.Get(s.URL + url)
	if err != nil {
		t.Fatal(err)
	}

	return resp, err
}

func (s *testServer) post(t *testing.T, url string, body []byte) (resp *http.Response, err error) {
	resp, err = http.Post(s.URL+url, "application/json", strings.NewReader(string(body)))

	if err != nil {
		t.Fatal(err)
	}

	return resp, err
}

func (s *testServer) put(t *testing.T, url string, body []byte) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, s.URL+url, strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	return resp, err
}

func (s *testServer) delete(t *testing.T, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodDelete, s.URL+url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	return resp, err
}

func (s *testServer) checkResponse(resp *http.Response, expectedStatusCode int, expectedBody string) error {
	if resp.StatusCode != expectedStatusCode {
		return errors.New("status code is not 200")
	}

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if strings.Contains(string(body), expectedBody) {
			return errors.New("body is not equal")
		}
	}

	return nil
}
