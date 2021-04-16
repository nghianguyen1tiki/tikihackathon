package httpfx

import "net/http"

func provideHTTPClient() (*http.Client, error) {
	return http.DefaultClient, nil
}
