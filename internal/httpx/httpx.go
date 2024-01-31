package httpx

import (
	"errors"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url必传")
	}

	res, err := http.Get(url)
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
