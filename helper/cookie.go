package helper

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, name string, value string, expired int64) {
	cookie := &http.Cookie{}
	cookie.Name = name
	cookie.Expires = time.Now().Add(time.Duration(expired) * time.Minute)
	cookie.Value = value
	cookie.Path = "/"

	http.SetCookie(w, cookie)

}

func GetCookie(request *http.Request, name string) string {
	cookie, err := request.Cookie(name)
	if err != nil {
		return ""
	} else {
		na := cookie.Value
		//fmt.Println("Hello %s", name)
		return na
	}
}
