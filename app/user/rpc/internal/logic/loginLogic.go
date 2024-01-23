package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"ticket-server/common/errs"

	"ticket-server/app/user/rpc/internal/svc"
	"ticket-server/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	username := in.Username
	password := in.Password
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)
	if err != nil {
		logx.Errorf("%+v", err)
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errs.UserNotExist
		}
		return nil, errs.MySQLInternalError
	}
	if user.Password != password {
		logx.Errorf("User[%s]，密码不匹配", username)
		return nil, errs.PasswordIncorrect
	}
	return &pb.LoginResp{
		Ok: true,
		Id: strconv.FormatInt(user.Id, 10),
	}, nil
}
