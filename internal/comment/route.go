package comment

import "glass-photo/internal/router"

func CommentRoute() {
	router.R.Get("/comments/{id}", getAllCommentsHandler)
	router.R.With(router.JWTMiddleware).Post("/comments/{id}", postCommentsHandler)
	router.R.With(router.JWTMiddleware).Delete("/comments/{id}", deleteCommentHandler)
}
