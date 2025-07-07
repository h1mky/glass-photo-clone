package user

import "glass-photo/internal/router"

func UserRouter() {
	router.R.Get("/user/{id}", getProfilePageHandler)
	router.R.With(router.JWTMiddleware).Get("/", GetMainPageInfo)
	router.R.With(router.JWTMiddleware).Patch("/user", updateUserHandler)
}
