package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

var (
	mockJsonData = []string{
		`{"email":"test1@gmail.com"}`,
		`{"email":"test2@gmail.com"}`,
	}
)

func TestCreateUser(t *testing.T) {
	e := echo.New()
	for _, data := range mockJsonData {
		req := httptest.NewRequest(http.MethodPost, "/createUser",
			strings.NewReader(data))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if err := CreateUser(c); err != nil {
			t.Error(err)
		}
	}

	req := httptest.NewRequest(http.MethodGet, "/getAllUsers", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := GetAllUsers(c); err != nil {
		t.Error(err)
	}

	var user []User
	if err := json.Unmarshal([]byte(rec.Body.String()), &user); err != nil {
		t.Error(err)
	}
	if user[0].ID != 1 && user[1].ID != 2 {
		t.Errorf("The result ID is not match. %#v", user)
	}
}
