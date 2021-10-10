package dto


// BookUpdateDTO is used by client when Put update book
type BookUpdateDTO struct {
	ID 			uint64	`json:"id" form :"id" binding:"required"`
	Tittle		string	`json:"tittle" form :"tittle" binding:"required"`
	Desciption	string	`json:"desciption" form :"desciption" binding:"required"`
	UserID		uint64	`json:"user_id, omitempty" form :"user_id, omitempty"`
}

// UserCreateDTO is used when create new book
type BookCreateDTO struct {
	Tittle		string	`json:"tittle" form :"tittle" binding:"required"`
	Desciption	string	`json:"desciption" form :"desciption" binding:"required"`
	UserID		uint64	`json:"user_id, omitempty" form :"user_id, omitempty"`
}


