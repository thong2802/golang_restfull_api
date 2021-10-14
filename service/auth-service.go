package service

import (
	"log"
	"github.com/mashingan/smapping"
	"github.com/thong2802/golang_api/dto"
	"github.com/thong2802/golang_api/entity"
	"github.com/thong2802/golang_api/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthSerivce is a contract about somthing that this service can do
type AuthSerivce interface {
	VerifyCredantial(email string, password string) interface{}
	CreatUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(mail string) bool
}

type authService struct {
	userRepository repository.UserRepository
}


// NewAuthSerivce creates a new instance of AuthSerivce
func NewAuthSerivce(userRap repository.UserRepository) AuthSerivce {
	return &authService{
		userRepository: userRap,
	}
}

func (service *authService) VerifyCredantial(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok :=res.(entity.User); ok {
		comparedPassword := comparedPassword(v.PassWord, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func(service *authService) CreatUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if  err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InserUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparedPassword(hashedPwd string, plainPassword []byte) bool{
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}