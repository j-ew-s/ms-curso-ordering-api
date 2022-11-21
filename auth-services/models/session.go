package authModels

import "github.com/j-ew-s/ms-curso-auth-grpc/auth-services"

type Session struct {
	Token           string
	TokenValidation *auth.TokenValidation
	Username        string
	Email           string
	Id              int64
	Name            string
}
