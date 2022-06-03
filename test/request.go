package test

import (
	"fmt"
	"github.com/injet-zhou/just-img-go-server/internal/router"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func PerformTestRequest(req *http.Request) ([]byte, error) {
	w := httptest.NewRecorder()
	r := router.RouteSetup()

	r.ServeHTTP(w, req)

	resp := w.Result()
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			fmt.Printf("close response body error: %v\n", closeErr)
		}
	}()

	data, err := ioutil.ReadAll(w.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
