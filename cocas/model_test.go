package cocas

import (
	"log"
	"testing"
)

func TestGetHealth(t *testing.T) {
	hea := GetHealth()

	if hea.Title != "IT works !" {
		log.Fatal("Not the good title")
	}
}

func TestGetProfile(t *testing.T) {
	//req := httptest.NewRequest(http.MethodGet, "/", nil)

	//GetProfile(req)
}
