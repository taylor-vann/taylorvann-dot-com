package cookies

import (
	"net/http"
)

const CookieDomain = "briantaylorvann.com"
const SessionCookieHeader = "briantaylorvann.com_session"
const InternalSessionCookieHeader = "briantaylorvann.com_internal_session"

const GuestSessionExpirationInSeconds = 60 * 60 * 24 * 3 

func GetSessionFromRequest(r *http.Request) (string, error) {
	sessionCookie, errSessionCookie := r.Cookie(SessionCookieHeader)
	if errSessionCookie != nil {
		return "", errSessionCookie
	}
	return sessionCookie.Value, errSessionCookie
}

func GetInternalSessionFromRequest(r *http.Request) (string, error) {
	sessionCookie, errSessionCookie := r.Cookie(InternalSessionCookieHeader)
	if errSessionCookie != nil {
		return "", errSessionCookie
	}
	return sessionCookie.Value, errSessionCookie
}

func AttachInternalGuestSessionToRequest(r *http.Request, session string) {
	r.AddCookie(&http.Cookie{
		Name:			InternalSessionCookieHeader,
		Value:		session,
		MaxAge:		GuestSessionExpirationInSeconds,
		Domain:   CookieDomain,
		Path:     "/",
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	})
}