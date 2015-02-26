package authen

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

var (
	s = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
)

func SetCookie(w http.ResponseWriter, value map[string]string) {
	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		fmt.Println(cookie)
		http.SetCookie(w, cookie)
	}
}

func GetCookie(r *http.Request, cookieName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		// fmt.Println(cookie.Value)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			fmt.Println(value)
		}
	}
}
