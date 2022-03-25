package handler

import (
	"net/http"
)

type Object struct {
	Filed string
}

func Test(w http.ResponseWriter, r *http.Request) Object {
	return Object{
		Filed: "test",
	}
}
