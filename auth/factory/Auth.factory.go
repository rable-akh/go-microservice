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

type Claims struct {
	Id   string `json:"id"`
	User string `json:"user"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func AuthLogin(requests requests.UserLoginRequest) (*pb.User, error) {

	user, err := GetUser(requests)
	if !err {
		fmt.Println("Error")
	}

	var result pb.User

	if utils.ComparePassword(requests.Password, user.Password) {

		var myClaims Claims

		mySignKey := []byte("akhgotesting")

		myClaims = Claims{
			User: user.Email,
			Id:   user.ID.Hex(),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
				Issuer:    "akh",
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				Subject:   "loginToken",
			},
		}

		// claims := &jwt.RegisteredClaims{
		// 	ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
		// 	Issuer:    "akh",
		// 	IssuedAt:  jwt.NewNumericDate(time.Now()),
		// 	Subject:   "loginToken",
		// 	ID:        user.Email,
		// }

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)

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

func CheckToken(request requests.CheckTokenRequest) (map[string]interface{}, error) {
	mySignKey := []byte("akhgotesting")
	claims := &Claims{}
	parse, err := jwt.ParseWithClaims(request.Token, claims, func(t *jwt.Token) (interface{}, error) {
		return mySignKey, nil
	})

	checkUser := ValidUser(requests.UserLoginRequest{UserName: claims.User})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return map[string]interface{}{
				"status":  "false",
				"message": "Token invalid",
			}, err
		}
		return map[string]interface{}{
			"status":  "false",
			"message": err.Error(),
		}, err
	}

	if checkUser {
		return map[string]interface{}{
			"status":  "true",
			"message": "Authorized",
		}, nil
	}

	if parse.Valid {
		return map[string]interface{}{
			"status":  "false",
			"message": "Unauthorized",
		}, err
	}

	return map[string]interface{}{
		"status":  "false",
		"message": "Unauthorized",
	}, err
}
