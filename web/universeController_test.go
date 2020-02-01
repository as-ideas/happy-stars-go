package web_test

import (
	"github.com/as-ideas/happy-stars-go/domain"
	"github.com/as-ideas/happy-stars-go/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const validAddUniverseCommand = `{
	"id": "someId",
	"name": "someName",
	"max_size": 12
}`

const invalidAddUniverseCommand = `{
	"id": "",
	"name": "",
	"max_size": 10
}`

func Test_UniverseController_AddUniverse(t *testing.T) {

	galaxy := test.GivenEmptyGalaxy()
	controller := test.GivenUniverseController(galaxy)

	payload := strings.NewReader(validAddUniverseCommand)
	req := httptest.NewRequest("POST", "/", payload)
	w := httptest.NewRecorder()

	controller.AddUniverse(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status %d, got '%v' instead", http.StatusCreated, resp.StatusCode)
	}

	universe, err := galaxy.GetUniverse("someId")
	if err != nil {
		t.Errorf("expected err to be nil, got '%s'", err)
	}

	expectedUniverse := domain.Universe{ID: "someId", Name: "someName", MaxSize: 12}

	if universe.ID != expectedUniverse.ID || universe.Name != expectedUniverse.Name || universe.MaxSize != expectedUniverse.MaxSize {
		t.Errorf("expected newly added universe to have same fields, \n"+
			"expected \t'%v'\ngot \t'%v", universe, expectedUniverse)
	}
}

func Test_UniverseController_AddUniverse_should_fail_with_BadRequest(t *testing.T) {

	galaxy := test.GivenEmptyGalaxy()
	controller := test.GivenUniverseController(galaxy)

	payload := strings.NewReader(invalidAddUniverseCommand)
	req := httptest.NewRequest("POST", "/", payload)
	w := httptest.NewRecorder()

	controller.AddUniverse(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got '%v' instead", http.StatusBadRequest, resp.StatusCode)
		t.Errorf("message is: %s", resp.Status)
	}

	_, err := galaxy.GetUniverse("someId")
	if err == nil {
		t.Errorf("expected err not to be nil")
	}
}

func Test_UniverseController_AddUniverse_should_fail_with_Conflict(t *testing.T) {

	galaxy := test.GivenEmptyGalaxy()
	galaxy.AddUniverse(domain.Universe{"someId", "someName", 12})
	controller := test.GivenUniverseController(galaxy)

	payload := strings.NewReader(validAddUniverseCommand)
	req := httptest.NewRequest("POST", "/", payload)
	w := httptest.NewRecorder()


	controller.AddUniverse(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusConflict {
		t.Errorf("expected status %d, got '%v' instead", http.StatusConflict, resp.StatusCode)
		t.Errorf("message is: %s", resp.Status)
	}
}
