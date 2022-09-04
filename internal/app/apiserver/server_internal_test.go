package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleIndex(t *testing.T) {
	s := New(NewConfig())
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	s.handleIndex().ServeHTTP(recorder, req)
	assert.Equal(t, recorder.Body.String(), "Hello world!")
}
