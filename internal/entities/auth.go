package entities

import authv1 "github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/proto/gen/auth"

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

func (r *RegisterRequest) ToGRPC() *authv1.RegisterRequest {
	return &authv1.RegisterRequest{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

func (r *LoginRequest) ToGRPC() *authv1.LoginRequest {
	return &authv1.LoginRequest{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}
