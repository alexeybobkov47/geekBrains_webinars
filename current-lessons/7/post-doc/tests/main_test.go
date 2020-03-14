package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers"

	"github.com/astaxie/beego"
)

func TestCreateListRequest(t *testing.T) {
	req := controllers.PostRequest{
		Name:        "New Test Req",
		Description: "New Test Descr",
	}

	body, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
	}

	r, _ := http.NewRequest("POST", "/v1/lists", bytes.NewReader(body))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("unexpected code: %v", w.Code)
	}
}
