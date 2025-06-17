package auth

import "glass-photo/internal/router"

func AuthRoute() {
	router.R.Post("/sign-up", createNewUserHandler)
	router.R.Post("/sign-in", SignInHandler)
}
