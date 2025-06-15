package post

type Post struct {
	Id          int    `json:"id" db:"id"`
	UserId      int    `json:"user_id" db:"user_id"`
	PostUrl     string `json:"post_url" db:"post_img"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}
