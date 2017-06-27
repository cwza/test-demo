package web

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"strconv"

	"github.com/cwza/test-demo/pkg/remote"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter_health(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "http status should be 200")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	require.NoError(t, err, "should not have error while get body")
	assert.Equal(t, "Hello, World", string(body), "resp should be Hello, World")
}

func TestRouter_seqGetValue(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	req := httptest.NewRequest("GET", "/seq/getValue", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "http status should be 200")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	require.NoError(t, err, "should not have error while get body")
	_, err = strconv.Atoi(string(body))
	assert.NoError(t, err, "result should be a int")
}

func TestRouter_seqReset(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	req := httptest.NewRequest("GET", "/seq/reset", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "http status should be 200")

	seq := remote.NewSequenceImpl()
	assert.Equal(t, 0, seq.GetValue())
}

func TestRouter_seqGetValueByStep(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := remote.NewSequenceImpl()
	defer seq.Reset()
	req := httptest.NewRequest("GET", "/seq/getValueByStep/3", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "http status should be 200")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	require.NoError(t, err, "should not have error while get body")
	value, err := strconv.Atoi(string(body))
	assert.NoError(t, err, "result should be a int")
	assert.Equal(t, 3, value)
}

func TestRouter_seqGetValueByStepError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := remote.NewSequenceImpl()
	defer seq.Reset()
	req := httptest.NewRequest("GET", "/seq/getValueByStep/xxx", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "http status should be 200")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	require.NoError(t, err, "should not have error while get body")
	assert.Equal(t, "step should be an integer", string(body), "should return error while step is xxx")
}
