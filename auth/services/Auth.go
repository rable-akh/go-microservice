package services

import (
	"context"
	"fmt"
	"log"
	"microservice/auth/factory"
	pb "microservice/auth/proto/microservice/auth"
	"microservice/auth/requests"
	"net/http"
	"strconv"
)

func (s *ServiceServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())

	data := requests.UserLoginRequest{
		UserName: in.User,
		Password: in.Pass,
	}

	user, err := factory.AuthLogin(data)

	var result pb.LoginResponse

	if err != nil {
		result = pb.LoginResponse{
			Status:  http.StatusNoContent,
			Message: err.Error(),
			Results: &pb.User{},
		}

		return &result, nil
	}

	result = pb.LoginResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Results: user,
	}

	return &result, nil
}

func (s *ServiceServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	log.Printf("Request: %v", in.ProtoReflect().Descriptor().FullName())

	data := requests.UserRequest{
		Name:     in.Name,
		Email:    in.Email,
		Phone:    in.Phone,
		Password: in.Password,
	}

	fmt.Println(data)

	results, err := factory.SaveUser(data)

	if !err {
		// log.Fatal(results)
		return &pb.SignUpResponse{
			Status:  200,
			Message: "Error",
			// Error:   "",
			Results: map[string]string{},
		}, nil
	}

	fmt.Println(results)

	return &pb.SignUpResponse{
		Status:  200,
		Message: "Success",
		Results: map[string]string{},
	}, nil
}

func (s *ServiceServer) GetUsers(ctx context.Context, in *pb.UsersRequest) (*pb.UsersResponse, error) {
	log.Printf("Request: %v", in.ProtoReflect().Descriptor().FullName())

	data := requests.PaginPara{
		Pages:   in.Page,
		PerPage: in.PerPage,
	}

	results, err := factory.GetUsers(data)

	if !err {
		log.Fatal(err)
	}

	var users []*pb.User

	for _, result := range results {
		users = append(users, &pb.User{
			Name:      result.Name,
			Email:     result.Email,
			Phone:     result.Phone,
			CreatedAt: result.CreatedAt.Time().Format("2006-01-02"),
		})
	}

	return &pb.UsersResponse{
		Status:  200,
		Message: "Success",
		Results: users,
	}, nil

}

func (s *ServiceServer) CheckToken(ctx context.Context, in *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	log.Printf("Request: %v", in.ProtoReflect().Descriptor().FullName())
	data := requests.CheckTokenRequest{
		Token: in.Token,
	}

	results, err := factory.CheckToken(data)

	status, _ := strconv.ParseBool(results["status"].(string))
	if err != nil {
		// log.Fatal(results)
		return &pb.CheckTokenResponse{
			Status:  status,
			Message: results["message"].(string),
		}, nil
	}

	return &pb.CheckTokenResponse{
		Status:  true,
		Message: "Success",
	}, nil
}
