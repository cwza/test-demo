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

	seq := remote.NewSequenceImpl()
	seq.Reset()
	req := httptest.NewRequest("GET", "/seq/getValue", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "http status should be 200")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	require.NoError(t, err, "should not have error while get body")
	value, err := strconv.Atoi(string(body))
	assert.NoError(t, err, "result should be a int")

	assert.Equal(t, 0, value)
}

func TestRouter_seqReset(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := remote.NewSequenceImpl()
	defer seq.Reset()

	seq.GetNext()
	req := httptest.NewRequest("GET", "/seq/reset", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "http status should be 200")

	assert.Equal(t, 0, seq.GetValue())
}

func TestRouter_seqGetNextByStep(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := remote.NewSequenceImpl()
	defer seq.Reset()

	seq.Reset()
	req := httptest.NewRequest("GET", "/seq/getNextByStep/3", nil)
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

func TestRouter_seqGetNextByStepError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := remote.NewSequenceImpl()
	defer seq.Reset()

	tcs := map[string]struct {
		A        string
		Expected int
	}{
		"minus":   {"-4", http.StatusNotFound},
		"not int": {"xxx", http.StatusNotFound},
		"zero":    {"0", http.StatusNotFound},
	}

	for msg, tc := range tcs {
		req := httptest.NewRequest("GET", "/seq/getNextByStep/"+tc.A, nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)
		actual := w.Result().StatusCode
		assert.Equal(t, tc.Expected, actual, msg)
	}
}
