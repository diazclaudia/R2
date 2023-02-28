package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ping_Ok(t *testing.T) {
	router := Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

func Test_SpiralFibonacci_Ok(t *testing.T) {
	router := Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/spiral?rows=2&cols=2", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"fibonacci_spiral\":[[0,1],[2,1]]}", w.Body.String())
}
