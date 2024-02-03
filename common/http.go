package common

import (
	"net/http"
)

func HttpDo(method string, url string, headers map[string]string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return client.Do(req)
}

func HttpGet(url string, headers map[string]string) (*http.Response, error) {
	return HttpDo("GET", url, headers)
}
