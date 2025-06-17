package post

import "glass-photo/internal/router"

func PostRouter() {
	router.R.Get("/post", GetAllPostHandler)
	router.R.Get("/post/{id}", GetPostByIDHandler)
	router.R.With(router.JWTMiddleware).Post("/post", createPostHandler)
}
