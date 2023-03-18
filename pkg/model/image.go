package model

type Image struct {
	Id      string `json:"id"`
	PostId  uint   `json:"postId"`
	Content string `json:"content"`
}
