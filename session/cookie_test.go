package session

import "testing"

func TestNewCookie(t *testing.T) {
	var cookie = NewCookie(nil)
	if cookie.HttpOnly != true {
		t.Error("Expected true but got", cookie.HttpOnly)
	}
}

func TestNewCookieExpiresToNull(t *testing.T) {
	var cookie = NewCookie(nil)
	if cookie.expires != nil {
		t.Error("Expected nil expires", cookie.expires)
	}
}

func TestShouldDefaultPathToSlash(t *testing.T) {
	var cookie = NewCookie(nil)
	if cookie.Path != "/" {
		t.Error("Expected / but got", cookie.Path)
	}
}

func TestShouldDefaultMaxAgeToNull(t *testing.T) {
	var cookie = NewCookie(nil)
	if cookie.GetMaxAge() != nil {
		t.Error("Expected nil but got", cookie.MaxAge)
	}
}
