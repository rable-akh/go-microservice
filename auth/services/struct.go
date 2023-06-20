package services

import (
	pb "microservice/auth/proto/microservice/auth"
)

type ServiceServer struct {
	pb.UnimplementedAuthServer
}
