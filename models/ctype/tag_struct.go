package ctype

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"Please Enter the Tag" structs:"title"`
}
