// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"ticket-server/app/user/rpc/internal/logic"
	"ticket-server/app/user/rpc/internal/svc"
	"ticket-server/app/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) ChangePassword(ctx context.Context, in *pb.ChangePasswordReq) (*pb.ChangePasswordResp, error) {
	l := logic.NewChangePasswordLogic(ctx, s.svcCtx)
	return l.ChangePassword(in)
}

func (s *UserServer) Info(ctx context.Context, in *pb.InfoReq) (*pb.InfoResp, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}
