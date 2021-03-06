package main

import (
	pb "github.com/heriyantoliu/Microservice/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var key = []byte("mySuperSecretKeyLo1")

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(token string) (*CustomClaims, error) {
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error){
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	}else{
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User)(string, error) {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt:15000,
			Issuer:"go.micro.srv.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
