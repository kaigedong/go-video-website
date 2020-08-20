package main

import (
	"net/http"

	"github.com/go-video-website/api/session"
)

var HEADER_FAILED_SESSION = "X-Session-Id" // X开头是自定义的Header
var HEADER_FIELED_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FAILED_SESSION)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELED_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELED_UNAME)
	if len(uname) == 0 {
		sendErrorResponse()
		return false
	}
	return true
}
