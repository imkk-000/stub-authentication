package atdd_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginPassed(t *testing.T) {
	expectedRespBody := []byte(`{"status":"ok"}`)
	reqBody := url.Values{}
	reqBody.Add("username", "imkk-000")
	reqBody.Add("password", "1mkkn@ja*")

	resp, err := http.PostForm(LOGIN_URL, reqBody)
	if err != nil {
		t.Fatal(err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, respBody)
	assert.Equal(t, expectedRespBody, respBody)
}
