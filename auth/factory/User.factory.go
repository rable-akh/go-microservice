package factory

import (
	"context"
	"fmt"
	"log"
	"microservice/auth/config"
	"microservice/auth/requests"
	"microservice/auth/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Password  string             `json:"password"`
	CreatedAt primitive.DateTime `json:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty"`
}

var userCollection *mongo.Collection = config.GetColl(config.DB, "users")

func GetUsers(requests requests.PaginPara) ([]User, bool) {

	filter := bson.D{}

	users, err := userCollection.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	var results []User

	if err = users.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	fmt.Println(results)
	return results, true

}

func GetUser(requests requests.UserLoginRequest) (User, bool) {

	filter := bson.D{{Key: "email", Value: requests.UserName}}

	var result User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return result, false
		}
		panic(err)
	}

	fmt.Println(result)

	return result, true
}

func SaveUser(requests requests.UserRequest) (interface{}, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pass, _ := utils.HashPassword(requests.Password)

	user := &User{
		Name:      requests.Name,
		Email:     requests.Email,
		Phone:     requests.Phone,
		Password:  pass,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return err, false
	}

	return result, true
}
