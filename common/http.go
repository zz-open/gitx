package common

import (
	"errors"
	"io"
	"net/http"
)

func HttpDo(method string, url string, headers map[string]string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url必传")
	}

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

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HttpGet(url string, headers map[string]string) ([]byte, error) {
	return HttpDo("GET", url, headers)
}

func HttpPost(url string, headers map[string]string) ([]byte, error) {
	return HttpDo("POST", url, headers)
}
