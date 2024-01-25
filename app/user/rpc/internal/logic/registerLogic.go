package logic

import (
	"context"
	"strconv"
	"ticket-server/app/user/rpc/internal/model"
	"ticket-server/common/errs"
	"time"

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
		Username:       username,
		Password:       password,
		LastChangeTime: time.Now(),
	})
	if err != nil {
		logx.Errorf("%+v", err)
		return nil, errs.MySQLRegisterError
	}
	var id int64
	if id, err = result.LastInsertId(); err != nil {
		logx.Errorf("%+v", err)
		return nil, errs.MySQLRegisterError
	}
	return &pb.RegisterResp{
		Ok: true,
		Id: strconv.FormatInt(id, 10),
	}, nil
}
