package middleware

import (
	"log"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thong2802/golang_api/helper"
	"github.com/thong2802/golang_api/service"
)

// AuthorizeJWT validate y
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) { 
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		response := helper.BuildErrorResponse("Failed tp process request", "No token found", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return 
	}
	token, err := jwtService.ValidateToken(authHeader) 
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims) 
		log.Println("Claim[user_id]: ", claims["user_id"])
		log.Panicln("Claim[issuer] :", claims["issuer"])
	}else {
		log.Println(err)
		response := helper.BuildErrorResponse("Token is not valid",  err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
  }
}