package session

import (
	"github.com/SamTV12345/cookie"
	"time"
)

type CookieParam struct {
	Path           *string
	HttpOnly       *bool
	OriginalMaxAge *int64
	Domain         *string
	SameSite       *string
	expires        *time.Time
	originalMaxAge *int64
	Secure         *bool
}

type Cookie struct {
	Path           string
	HttpOnly       bool
	OriginalMaxAge *int64
	Domain         string
	SameSite       string
	expires        *time.Time
	originalMaxAge *int64
	Secure         bool
}

func NewCookie(options interface{}) Cookie {
	var sessCookie = Cookie{
		Path:     "/",
		HttpOnly: true,
	}
	sessCookie.SetMaxAge(nil)

	if options != nil && options.(*CookieParam) != nil {
		var optionsCookie = options.(*CookieParam)
		if optionsCookie.Path != nil {
			sessCookie.Path = *optionsCookie.Path
		}
		if optionsCookie.HttpOnly != nil {
			sessCookie.HttpOnly = *optionsCookie.HttpOnly
		}
		if optionsCookie.expires != nil {
			sessCookie.expires = optionsCookie.expires
		}

		if optionsCookie.Domain != nil {
			sessCookie.Domain = *optionsCookie.Domain
		}
		if optionsCookie.SameSite != nil {
			sessCookie.SameSite = *optionsCookie.SameSite
		}
		if optionsCookie.Secure != nil {
			sessCookie.Secure = *optionsCookie.Secure
		}
	}

	if sessCookie.OriginalMaxAge == nil {
		var newMaxAge = sessCookie.GetMaxAge()
		sessCookie.OriginalMaxAge = newMaxAge
	}

	return sessCookie
}

func (c *Cookie) SetExpires(timeOfExpiracy time.Time) {
	c.expires = &timeOfExpiracy
	var newMaxAge = c.GetMaxAge()
	c.originalMaxAge = newMaxAge
}

// GetExpires returns the expiration time of the cookie
func (c *Cookie) GetExpires() *time.Time {
	return c.expires
}

func (c *Cookie) GetMaxAge() *int64 {
	if c.expires == nil {
		return nil
	}
	var result = c.expires.UnixMilli()
	return &result
}

func (c *Cookie) SetMaxAge(ms *int64) {
	if ms == nil {
		c.expires = nil
		return
	}
	var newExpires = time.UnixMilli(time.Now().UnixMilli() + *ms)
	c.expires = &newExpires
}

func (c *Cookie) GetData() Cookie {
	return Cookie{
		OriginalMaxAge: c.originalMaxAge,
		expires:        c.expires,
		Secure:         c.Secure,
		HttpOnly:       c.HttpOnly,
		Domain:         c.Domain,
		Path:           c.Path,
		SameSite:       c.SameSite,
	}
}

func (c *Cookie) Serialize(name, val string) (*string, error) {
	return cookie.Serialize(name, val, nil)
}
