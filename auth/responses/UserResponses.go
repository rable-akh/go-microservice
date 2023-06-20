package responses

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	CreatedAt primitive.DateTime `json:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at"`
}

type UsersResponse struct {
	Status  string
	Message string
	Results []UserResponse
}
