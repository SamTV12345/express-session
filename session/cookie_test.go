package session

import (
	"testing"
	"time"
)

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
		t.Error("Expected nil but got", cookie.GetMaxAge())
	}
}

func TestShouldCreateCookieWithOptions(t *testing.T) {
	var result = NewCookie(&CookieParam{})
	if result.GetMaxAge() != nil {
		t.Error("Expected nil but got", result.GetMaxAge())
	}
}

func TestShouldSetExpires(t *testing.T) {
	var expires = time.UnixMilli(time.Now().UnixMilli() + 60000)
	var cookie = NewCookie(&CookieParam{
		expires: &expires,
	})
	if cookie.expires.UnixMilli() != expires.UnixMilli() {
		t.Error("Expected", expires, "but got", cookie.expires)
	}
}

func TestShouldSetMaxAge(t *testing.T) {
	expires := time.UnixMilli(time.Now().UnixMilli() + 60000)
	cookie := NewCookie(&CookieParam{
		expires: &expires,
	})

	if !(expires.UnixMilli()-time.Now().UnixMilli()-1000 <= *cookie.GetMaxAge()) {
		t.Error("Expected", expires, "but got", cookie.expires)
	}

	if !(expires.UnixMilli()-time.Now().UnixMilli()+1000 <= *cookie.GetMaxAge()) {
		t.Error("Expected", expires, "but got", cookie.expires)
	}
}

func TestShouldSetHttOnly(t *testing.T) {
	var httpOnly = false
	var cookie = NewCookie(&CookieParam{
		HttpOnly: &httpOnly,
	})

	if cookie.HttpOnly != false {
		t.Error("Expected false but got", cookie.HttpOnly)
	}
}
