package web_test

import (
	"encoding/json"
	"github.com/as-ideas/happy-stars-go/test"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_InfoController_ServeColorValues(t *testing.T) {

	galaxy := test.GivenEmptyGalaxy()
	controller := test.GivenInfoController(galaxy)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	controller.ServeColorValues(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got '%v' instead", resp.StatusCode)
	}

	expectedContentType := "application/json"
	actualContentType := resp.Header.Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("expected '%s', got '%v' instead", expectedContentType, actualContentType)
	}

	var colors []string
	json.Unmarshal(body, &colors)
	if len(colors) < 1 {
		t.Errorf("expected colors, got '%d' instead", len(colors))
	}
}
