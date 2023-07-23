package models

type Comment struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	PostID    string `json:"postId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}