package logic

import (
	"context"
	"errors"
	"fmt"
	"ticket-server/app/user/rpc/internal/model"

	"ticket-server/app/user/rpc/internal/svc"
	"ticket-server/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	username := in.Username
	password := in.Password
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil, errors.New("注册失败")
	}
	if c, _ := result.RowsAffected(); c != 1 {
		return nil, errors.New("注册失败")
	}
	return &pb.RegisterResp{
		Ok: true,
		Id: "注册成功",
	}, nil
}
