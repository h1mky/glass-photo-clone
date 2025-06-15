package user

import "glass-photo/internal/router"

func UserRouter() {
	router.R.Get("/user/{id}", getProfilePageHandler)
}
