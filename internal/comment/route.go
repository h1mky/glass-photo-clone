package comment

import "glass-photo/internal/router"

func CommentRoute() {
	router.R.Get("/comments/{id}", getAllCommentsHandler)
}
