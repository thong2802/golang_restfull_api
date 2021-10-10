package dto


// UserUpdateDTO is used by client when Put update profile
type UserUpdateDTO  struct{
	ID 			uint64	`json:"id" form :"id" binding:"required"`
	Name		string	`json:"name" form :"name" binding:"required"`
	Email		string	`json:"email" form :"email" binding:"required" validate:"email"`	
	Password	string	`json:"password, omitempty" form:"password, omitempty" validate: "min:6"`
}

// UserCreateDTO is used when create new user
type UserCreateDTO  struct{
	ID 			uint64	`json:"id" form :"id" binding:"required"`
	Name		string	`json:"name" form :"name" binding:"required"`
	Email		string	`json:"email" form :"email" binding:"required" validate:"email"`	
	Password	string	`json:"password, omitempty" form:"password, omitempty" validate: "min:6" binding:"required"`
}


