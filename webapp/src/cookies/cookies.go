package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New(config.HASH_KEY, config.BLOCK_KEY)
}

func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	dataEncoded, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    dataEncoded,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	c, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)

	if err := s.Decode("data", c.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}
