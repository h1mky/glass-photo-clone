package post

type MainPagePost struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	PostImg  string `db:"post_img" json:"post_img"`
	Username string `db:"username" json:"username"`
}
