package repository

import (
	"log"

	"github.com/thong2802/golang_api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what UserRepository can to do db
type UserRepository interface {
	InserUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func(db *userConnection) InserUser(user entity.User) entity.User {
	user.PassWord = hashAndsalt([]byte(user.PassWord))
	db.connection.Save(&user)
	return user
}

func(db *userConnection) UpdateUser(user entity.User) entity.User {
	user.PassWord = hashAndsalt([]byte(user.PassWord))
	db.connection.Save(&user)
	return user
}
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res :=db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil  {
		return user
	}
	return nil
}

func(db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func(db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func(db *userConnection) ProfileUser( userID string) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}


func hashAndsalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if  err != nil {
		log.Println(err)
		panic("Fail to hash password")
	}
	return string(hash)
}