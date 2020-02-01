package web

import (
	"encoding/json"
	"github.com/as-ideas/happy-stars-go/domain"
	"github.com/as-ideas/happy-stars-go/usecases"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInfoController_ServeColorValues(t *testing.T) {

	controller := InfoController{InfoUsecase: usecases.InfoUsecase{Galaxy: &domain.Galaxy{}}}

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
