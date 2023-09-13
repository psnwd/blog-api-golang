package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID          string   `json:"id"`
	UserID      string   `json:"userId"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type Comment struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	PostID    string `json:"postId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}