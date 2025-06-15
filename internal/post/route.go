package post

import "glass-photo/internal/router"

func PostRouter() {
	router.R.Get("/posts", getAllPostHandler)
}
