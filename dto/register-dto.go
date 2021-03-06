package dto

// RegisterDTO is used when client post from / register url
type RegisterDTO struct {
	Name		string	`json:"name" form: "name" binding:"required" vaidate:"min:1"`
	Email		string 	`json:"email" form:"email" binding:"required" vaidate:"email"`
	Password	string	`json:"password" form:"password" binding:"required" validate: "min:6"`
}