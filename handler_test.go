package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {

	mockedUserStore := &mockUserStore{}

	userService := NewUserHandler(mockedUserStore)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(userService.registerUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected a status code of %v but got %v", http.StatusOK, status)
	}
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByID(id int) (*User, error) {
	return nil, nil
}

func (m *mockUserStore) DeleteUserByID(id int) error {
	return nil
}
