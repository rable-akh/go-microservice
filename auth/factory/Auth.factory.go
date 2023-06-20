package factory

import (
	"fmt"
	"log"
	pb "microservice/auth/proto/microservice/auth"
	"microservice/auth/requests"
	"microservice/auth/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func AuthLogin(requests requests.UserLoginRequest) (*pb.User, error) {

	user, err := GetUser(requests)
	if !err {
		fmt.Println("Error")
	}

	var result pb.User

	if utils.ComparePassword(requests.Password, user.Password) {

		mySignKey := []byte("akhgotesting")

		claims := &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			Issuer:    "akh",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "loginToken",
			ID:        user.Email,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		atoken, error := token.SignedString(mySignKey)

		if error != nil {
			log.Default().Fatalln(error)
		}

		result = pb.User{
			XId:       user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Atoken:    atoken,
			CreatedAt: user.CreatedAt.Time().Format("2006-01-02"),
		}
		return &result, nil
	} else {
		return &pb.User{}, utils.ErrorPasswordVerify
	}
}
