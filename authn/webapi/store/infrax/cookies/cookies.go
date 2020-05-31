package cookies

import (
	"net/http"
)

const CookieDomain = ".briantaylorvann.com"
const SessionCookieHeader = "briantaylorvann.com_session"

const GuestSessionExpirationInSeconds = 60 * 60 * 24 * 3 

func CreateGuestSessionCookie(session string) *http.Cookie {
	return &http.Cookie{
		Name:			SessionCookieHeader,
		Value:		session,
		MaxAge:		GuestSessionExpirationInSeconds,
		Domain:   CookieDomain,
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	}
}

func GetSessionCookie(r http.Request) (*http.Cookie, error) {
	return r.Cookie(SessionCookieHeader)
}