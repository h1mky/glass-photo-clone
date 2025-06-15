package post

import "glass-photo/internal/router"

func PostRouter() {
	router.R.Get("/post", getAllPostHandler)
	router.R.Get("/post/{id}", getPostByIDHandler)
}
