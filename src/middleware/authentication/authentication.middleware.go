package AuthentcationMiddleware

import (
	"strings"
	"net/http"
	"github.com/myriadeinc/pickaxe/src/util/config"
)

type Decorator func(http.Handler) http.Handler

func Authenticate() Decorator {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			authArr := strings.Fields(req.Header.Get("Authorization"))
			if (2 == len(authArr) && "shared_secret" == authArr[0] && validateSharedSecret(authArr[1])){
				h.ServeHTTP(res, req)
				return
			}
			http.Error(res, "Not Found", 404)
		})
	}
}

func validateSharedSecret(secret string) bool {
	expected := ConfigUtil.Get("service.shared_secret")
	return (expected == secret)
}