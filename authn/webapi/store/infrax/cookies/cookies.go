package cookies

import (
	"net/http"
)

const CookieDomain = "briantaylorvann.com"
const SessionCookieHeader = "briantaylorvann.com_session"
const InternalSessionCookieHeader = "briantaylorvann.com_internal_session"

const GuestSessionExpirationInSeconds = 60 * 60 * 24 * 3 

func GetInternalGuestSessionCookie(r *http.Request, session string) *http.Cookie {
	return http.Cookie{
		Name:			InternalSessionCookieHeader,
		Value:		session,
		MaxAge:		GuestSessionExpirationInSeconds,
		Domain:   CookieDomain,
		Path:     "/",
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	}
}