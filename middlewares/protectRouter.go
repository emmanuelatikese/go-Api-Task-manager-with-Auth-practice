package middlewares

import (
	"net/http"

	jwtFunc "api.task/jwt"
	apiUtils "api.task/utils"
)

func ProtectRoutes(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	cookies, err := r.Cookie("jwt_token")
	if err != nil {
		apiUtils.JsonResponse("Not authorized", w, http.StatusUnauthorized)
		return
	}
	value := cookies.Value
	if value == "" {
		apiUtils.JsonResponse("Not authorized", w, http.StatusUnauthorized)
		return
	}

	id, err := jwtFunc.Verify(w, r, value)
	if err != nil {
		apiUtils.JsonResponse("Not authorized", w, http.StatusUnauthorized)
		return
	}
	if id == "" || id == nil {
		apiUtils.JsonResponse("Not authorized", w, http.StatusUnauthorized)
		return
	}
	next(w,r)
}