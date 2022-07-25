package auth

import (
	"net/http"
	"testing"
)

// VerifyCookieHelper is a helper function used in unit tests.
//
// It parses the cookieValue and checks its validity using the provided auth.
// If the cookie cannot be parsed or the validation fails, it notifies t of the error.
func VerifyCookieHelper(t *testing.T, auth *Authenticator, cookieValue string) {
	// Create request object that a client would receive and check cookie
	clientRequest := &http.Request{}
	clientRequest.Header = http.Header{}
	clientRequest.Header.Set("Cookie", cookieValue)

	// Parse cookie and check validity
	if cookie, err := clientRequest.Cookie(auth.CookieName()); err == nil {
		if _, err := auth.Verify(cookie.Value); err != nil {
			t.Error(err)
		}
	} else {
		t.Error(err)
	}
}
