package helper

import (
	"net/http"
	"strings"
	"time"
)

func SetCookie(w http.ResponseWriter, name string, value string, expired int64) {
	cookie := &http.Cookie{}
	cookie.Name = strings.ToUpper(name)
	cookie.Expires = time.Now().Add(time.Duration(expired) * time.Minute)
	cookie.Value = value
	cookie.Path = "/"

	http.SetCookie(w, cookie)

}

func GetCookie(request *http.Request, name string) string {
	cookie, err := request.Cookie(strings.ToUpper(name))
	if err != nil {
		return ""
	} else {
		na := cookie.Value
		//fmt.Println("Hello %s", name)
		return na
	}
}
