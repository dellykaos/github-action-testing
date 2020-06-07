package healthz_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dellykaos/github-action-testing/handler/healthz"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(healthz.New())
}

func TestHealthz(t *testing.T) {
	assert := assert.New(t)

	h := healthz.New()

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/healthz", nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Healthz(w, r, httprouter.Params{})
	})
	handler.ServeHTTP(resp, req)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(200, resp.Code)
	assert.Equal("ok", string(body))
}
