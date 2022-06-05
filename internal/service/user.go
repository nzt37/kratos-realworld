package service

import (
	"context"

	v1 "kartos-realworld/api/realworld/v1"

)

func (s *RealworldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User:&v1.UserReply_User{
			Username:"boom",
		},
	},nil
}