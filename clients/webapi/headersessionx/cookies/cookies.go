package cookies

import (
	"net/http"
)

const CookieDomain = "briantaylorvann.com"
const UserSessionCookieHeader = "briantaylorvann.com_session"
const GuestSessionExpirationInSeconds = 60 * 60 * 24 * 3 

func AttachGuestSession(w http.ResponseWriter, session string) {
	guestSessionCookie := http.Cookie{
		Name:			UserSessionCookieHeader,
		Value:		session,
		MaxAge:		GuestSessionExpirationInSeconds,
		Domain:   CookieDomain,
		Path:     "/",
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	}
	
	http.SetCookie(w, &guestSessionCookie)
}