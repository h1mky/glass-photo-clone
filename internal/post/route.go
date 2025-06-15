package post

import "glass-photo/internal/router"

func PostRouter() {
	router.R.Get("/post", getAllPostHandler)
}
